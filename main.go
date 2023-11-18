package main

import (
	"Gateway/conf"
	"Gateway/internal/controllers"
	"fmt"
)

func main() {
	// server模式运行
	port := conf.Get().Service.Port
	_ = controllers.SetRouters().Run(fmt.Sprintf(":%d", port))
}
