package common

import (
	"strconv"
	"time"
)

const (
	nodeBits uint8 = 5 // 节点 ID 的位数
	seqBits  uint8 = 7 // 序列号的位数
	nodeMax  int64 = -1 ^ (-1 << nodeBits)
	seqMask  int64 = -1 ^ (-1 << seqBits)
)

type Worker struct {
	lastTimestamp int64
	node          int64
	sequence      int64
}

func NewWorker(workerId int64) (*Worker, error) {
	return &Worker{
		lastTimestamp: 0,
		node:          workerId,
		sequence:      0,
	}, nil
}

func (w *Worker) GetId() string {
	timestamp := time.Now().UnixNano() / 1000000 // 当前时间的毫秒数

	if w.lastTimestamp == timestamp {
		w.sequence = (w.sequence + 1) & seqMask
		if w.sequence == 0 {
			for timestamp <= w.lastTimestamp {
				timestamp = time.Now().UnixNano() / 1000000
			}
		}
	} else {
		w.sequence = 0
	}
	w.lastTimestamp = timestamp
	id := ((timestamp - 1288834974657) << (nodeBits + seqBits)) | (w.node << seqBits) | w.sequence

	return strconv.FormatInt(id, 10)
}

func GenerateId() string {
	node, err := NewWorker(1)
	if err != nil {
		panic(err)
	}

	return node.GetId()
}
