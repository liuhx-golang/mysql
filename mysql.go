package mysql

import (
	"fmt"
	"time"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB = &gorm.DB{}

func Init2(user string, password string, host string, db string, connMaxLifetime int, maxIdleConns int, mxOpenConns int) {
	link := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, db)
	DB, err := gorm.Open(mysql.New(mysql.Config{
		DSN: link, // DSN data source name
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("connect mysql error ="+err.Error())
	}
	sqlDB, err := DB.DB()
	if err != nil {
		panic("connect mysql error ="+err.Error())
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(maxIdleConns)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(mxOpenConns)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(connMaxLifetime * time.Millisecond)
}
