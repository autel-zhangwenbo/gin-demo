package dao

import (
	"fmt"
	"gin-demo/internal/configs"
	"gin-demo/internal/models/article"
	log "github.com/sirupsen/logrus"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB



func Init() {
	dbConf := configs.Cfg.Database
	fmt.Println(dbConf)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&loc=Asia%%2FShanghai&parseTime=true", dbConf.User, dbConf.Password, dbConf.Host, dbConf.DbName)
	var err error
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("init ota mysql database failed", zap.Error(err))
	}

	if err = Db.AutoMigrate(&article.Article{}); err != nil {
		log.Fatal("database auto migrate failed", zap.Error(err))
	}
}

type PageParam struct {
	Page int
	Size int
}