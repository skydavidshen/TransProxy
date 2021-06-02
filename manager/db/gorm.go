package db

import (
	"TransProxy/manager"
	"bytes"
	"fmt"
	mysqldriver "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Gorm() (*gorm.DB, error) {
	dbConfig := manager.TP_SERVER_CONFIG.System.Db
	switch dbConfig {
	case "mysql":
		return mysql()
	default:
		return mysql()
	}
}

func mysql() (*gorm.DB, error) {
	config := mysqldriver.Config{
		DSN: getDsn(),         // DSN data source name
		DefaultStringSize: 180,
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}

	if db, err := gorm.Open(mysqldriver.New(config)); err != nil {
		return nil, err
	} else {
		sql, errDB := db.DB()
		if errDB != nil {
			return nil, errDB
		}
		sql.SetMaxIdleConns(manager.TP_SERVER_CONFIG.DB.Mysql.MaxIdleConn)
		sql.SetMaxOpenConns(manager.TP_SERVER_CONFIG.DB.Mysql.MaxOpenConn)
		return db, nil
	}
}

func getDsn() string {
	var buffer bytes.Buffer
	buffer.WriteString(manager.TP_SERVER_CONFIG.DB.Mysql.Username)
	buffer.WriteString(":")
	buffer.WriteString(manager.TP_SERVER_CONFIG.DB.Mysql.Password)
	buffer.WriteString("@tcp(")
	buffer.WriteString(manager.TP_SERVER_CONFIG.DB.Mysql.Host)
	buffer.WriteString(")/")
	buffer.WriteString(manager.TP_SERVER_CONFIG.DB.Mysql.DBName)
	buffer.WriteString("?")
	buffer.WriteString(manager.TP_SERVER_CONFIG.DB.Mysql.Option)
	return buffer.String()
}

// Closes the database and prevents new queries from starting.
func Close() {
	sqlDB, _ := manager.TP_DB.DB()
	err := sqlDB.Close()
	if err !=nil {
		panic(fmt.Errorf("Close sql connection Failed err: %s \n", err))
	}
}