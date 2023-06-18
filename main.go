package main

import (
	"recommendation/initial"
)

func main() {
	initial.Init()
	r := initial.Routers()
	panic(r.Run(":9090"))

}
