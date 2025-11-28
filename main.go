package main

import (
	"fmt"

	"github.com/LikkasT/GO-BLOG/biz/db"
	"github.com/LikkasT/GO-BLOG/biz/middle_ware"
	"github.com/LikkasT/GO-BLOG/biz/models"
	"github.com/LikkasT/GO-BLOG/biz/router"
)

func main() {
	router_viper, err := middle_ware.ConfigInit("config", ".", "yaml")
	if err != nil {
		fmt.Printf("error Init viper, %s", err)
		return
	}
	logger_level := middle_ware.Logger_level_map[router_viper.GetString("logger_level")]
	logger_file_name := router_viper.GetString("logger_file_name")
	router_logger, err := middle_ware.LoggerInit(logger_level, logger_file_name)
	if err != nil {
		fmt.Printf("error Init logger, %s", err)
		return
	}
	router_db, err := db.DbInit(router_viper, router_logger)
	if err != nil {
		fmt.Printf("error Init db, %s", err)
		return
	}
	err = router_db.AutoMigrate(&models.Users{})
	if err != nil {
		fmt.Printf("error migrate db, %s", err)
		return
	}
	routerDeps := router.RouterDependencies(router_viper, router_logger, router_db)
	routerDeps.RouterInit()

}
