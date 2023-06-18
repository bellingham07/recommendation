package RedisUtil

import (
	"GoRedisLearn/util"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type LockForRedis struct {
	// redis 客户端
	store *redis.Client
	// 超时时间
	seconds time.Duration
	// 续期时间
	renewal time.Duration
	// 锁key
	keys string
	// 用户解锁通知通道
	unlockCh chan struct{}
	// 雪花算法生成id 防止误删
	value string
}

func NewLockForRedis(keys string) *LockForRedis {
	return &LockForRedis{
		store:    GetClient(),
		seconds:  20,
		renewal:  8,
		keys:     "order_no:" + keys, // TODO 加前缀
		unlockCh: make(chan struct{}, 0),
		value:    util.GenSnowFlakeId(),
	}
}

func (lock *LockForRedis) Lock() (bool, error) {
	var resp *redis.BoolCmd
	_, rerr := lock.store.Ping(ctx).Result()
	if rerr != nil {
		return false, rerr
	}
	for {
		resp = lock.store.SetNX(ctx, lock.keys, lock.value, time.Second*lock.seconds) // 返回执行结果
		lockSuccess, err := resp.Result()
		if err == nil || lockSuccess {
			go lock.watchDog()
			return true, err
		} else {
			// 抢锁失败 回到客户端
			// TODO 这里应该可以让他等一段时间 在进行加载 就是看在等待时间内 前一个锁能不能被释放
			return false, err // TODO 自己写error信息
		}
	}
}

func (lock *LockForRedis) Unlock() {
	script := redis.NewScript(`
	if redis.call('get',KEYS[1])==ARGV[1]
	then
		return redis.call('del',KEYS[1])
	else
		return 0
	end
  `)
	client := RedisUtil.Client
	resp := script.Run(ctx, client, []string{lock.keys}, lock.value)
	if result, err := resp.Result(); err != nil || result == 0 {
		fmt.Println("unlock failed:", err)
	} else {
		// 删锁成功 通知看门狗退出
		lock.unlockCh <- struct{}{}
	}
}

func (lock *LockForRedis) watchDog() {
	expTicker := time.NewTicker(time.Second * lock.renewal)
	client := RedisUtil.Client

	// 确认锁与锁续期打包原子化
	script := redis.NewScript(`
	if redis.call('get',KEYS[1])==ARGV[1]
	then
		return redis.call('expire',KEYS[1],ARGV[2])
	else
		return 0
	end
   `)
	for {
		select {
		case <-expTicker.C:
			resp := script.Run(ctx, client, []string{lock.keys}, lock.value, 10)
			if result, err := resp.Result(); err != nil || result == int64(0) {
				// 续期失败
				panic(err)
			}
		case <-lock.unlockCh: // 任务完成后用户解锁 通知看门狗退出
			return
		}
	}
}
