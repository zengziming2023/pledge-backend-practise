package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"pledge-backend-practise/config"
	"pledge-backend-practise/log"
	"time"
)

func InitMySql() {
	log.Logger.Info("init mysql")
	mysqlConf := config.Config.Mysql

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConf.UserName,
		mysqlConf.Password,
		mysqlConf.Address,
		mysqlConf.Port,
		mysqlConf.DbName)

	dialector := mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	})

	db, err := gorm.Open(dialector, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 关闭复数表(表名后缀加上了s)
		},
		SkipDefaultTransaction: true,
	})
	if err != nil {
		//panic("failed to connect database: " + err.Error())
		log.Logger.Error("failed to connect database: " + err.Error())
		return
	}

	_ = db.Callback().Create().After("gorm:after_create").Register("after_create", After)
	_ = db.Callback().Query().After("gorm:after_query").Register("after_query", After)
	_ = db.Callback().Delete().After("gorm:after_delete").Register("after_delete", After)
	_ = db.Callback().Update().After("gorm:after_update").Register("after_update", After)
	_ = db.Callback().Row().After("gorm:row").Register("after_row", After)
	_ = db.Callback().Raw().After("gorm:raw").Register("after_raw", After)

	sqlDB, err := db.DB()
	if err != nil {
		log.Logger.Error("db.Db error: " + err.Error())
		return
	}

	sqlDB.SetMaxIdleConns(mysqlConf.MaxIdleConns) // 空闲连接数   默认最大2个空闲连接数  使用默认值即可
	sqlDB.SetMaxOpenConns(mysqlConf.MaxOpenConns) // 最大连接数   默认0是无限制的  使用默认值即可
	sqlDB.SetConnMaxLifetime(time.Duration(mysqlConf.MaxLifeTime) * time.Second)

	MySql = db
}

func After(db *gorm.DB) {
	sql := db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)
	log.Logger.Info("explain sql: " + sql)
}
