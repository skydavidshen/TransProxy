package db

import (
	"bytes"
	"com.pippishen/trans-proxy/manager"
	"fmt"
	"gorm.io/gorm"
	mysqldriver "gorm.io/driver/mysql"
)

func Gorm() *gorm.DB {
	dbConfig := manager.TP_CONFIG.Get("system.db").(string)
	switch dbConfig {
	case "mysql":
		return mysql()
	default:
		return mysql()
	}
}

func mysql() *gorm.DB {
	mysqlConfig := manager.TP_CONFIG.Get("db.mysql").(map[string]interface{})
	config := mysqldriver.Config{
		DSN: getDsn(mysqlConfig),         // DSN data source name
		DefaultStringSize: 180,
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysqldriver.New(config)); err != nil {
		return nil
	} else {
		sql, _ := db.DB()
		sql.SetMaxIdleConns(mysqlConfig["max-idle-conns"].(int))
		sql.SetMaxOpenConns(mysqlConfig["max-open-conns"].(int))
		return db
	}
}

func getDsn(mysqlConfig map[string]interface{}) string {
	var buffer bytes.Buffer
	buffer.WriteString(mysqlConfig["username"].(string))
	buffer.WriteString(":")
	buffer.WriteString(mysqlConfig["password"].(string))
	buffer.WriteString("@tcp(")
	buffer.WriteString(mysqlConfig["host"].(string))
	buffer.WriteString(")/")
	buffer.WriteString(mysqlConfig["db-name"].(string))
	buffer.WriteString("?")
	buffer.WriteString(mysqlConfig["option"].(string))
	return buffer.String()
}

// closes the database and prevents new queries from starting.
func Close() {
	sqlDB, _ := manager.TP_DB.DB()
	err := sqlDB.Close()
	if err !=nil {
		panic(fmt.Errorf("Close sql connection Failed err: %s \n", err))
	}
}