package main

import (
	"fmt"

	"github.com/rhdedgar/pod-logger/routers"
)

func main() {
	fmt.Println("Pod-logger V0.8")

	e := routers.Routers
	e.Logger.Info(e.Start(":8080"))
}
