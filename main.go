package main

import (
	"fmt"

	"github.com/LikkasT/GO-BLOG/biz/middle_ware"
	"github.com/LikkasT/GO-BLOG/biz/router"
)

func main() {
	router_viper, err := middle_ware.ConfigInit("config", ".", "yaml")
	if err != nil {
		fmt.Printf("error reading config file, %s", err)
		return
	}
	router := router.RouterDependencies(router_viper)
	router.RouterInit()

}
