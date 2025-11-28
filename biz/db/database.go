package db

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DbInit(cfg *viper.Viper, logger *logrus.Logger) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=Local",
		cfg.GetString("MySQL.DBUser"),
		cfg.GetString("MySQL.DBPassword"),
		cfg.GetString("MySQL.DBHost"),
		cfg.GetInt("MySQL.DBPort"),
		cfg.GetString("MySQL.DBName"),
		cfg.GetString("MySQL.DBCharset"),
		cfg.GetBool("MySQL.DBIfParseTime"))
	logger.Infof("database connect message %s", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, fmt.Errorf("error open db, %s", err)
	}
	return db, nil
}
