package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	log "github.com/liuhx-golang/log4go"
	"time"
)

//MySQLClient s
var MySQLClient = &gorm.DB{}
var connErr interface{}

//Init mysql
func Init(user string, password string, host string, db string, maxOpenConns time.Duration, maxIdleConns int, mxOpenConns int) {
	link := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, db)
	MySQLClient, connErr = gorm.Open(
		"mysql",
		link,
	)
	if connErr != nil {
		log.GetLogger().Error("mysql链接报错", connErr)
	}
	MySQLClient.LogMode(true)
	MySQLClient.DB().SetConnMaxLifetime(maxOpenConns * time.Millisecond)
	MySQLClient.DB().SetMaxIdleConns(maxIdleConns)
	MySQLClient.DB().SetMaxOpenConns(mxOpenConns)

}
