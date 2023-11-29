package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"newframework/pkg/config"
	myLog "newframework/pkg/log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var Database *gorm.DB

// Init 初始化数据库
func InitDB() {
	createDatabase()

	var err error
	c := config.GlobalConfig.DBConfig

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		c.Host, c.Username, c.Password, c.Dbname, c.Port,
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             5 * time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Error,    // 日志级别
			IgnoreRecordNotFoundError: true,            // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,           // 禁用彩色打印
		},
	)
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction:                   true, //关闭默认事务
		DisableForeignKeyConstraintWhenMigrating: true, //迁移时 禁用自动创建外键约束
		Logger:                                   newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 关闭表名默认复数形式
		},
	})
	if err != nil {
		myLog.Log().Fatalln("数据库连接失败：", err.Error())
		return
	}
	sqlDB, err := Database.DB()
	if err != nil {
		myLog.Log().Fatalln("数据库连接失败：", err.Error())
		return
	}

	// SetMaxOpenConns 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(c.MaximumPoolSize)
	// SetConnMaxLifetime 设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)
	//连接池最大空闲数
	sqlDB.SetMaxIdleConns(c.MaximumIdleSize)
}

// GetDB 实例
func GetDB(c context.Context) *gorm.DB {
	return Database.WithContext(c)
}

func createDatabase() {
	c := config.GlobalConfig.DBConfig

	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		c.Host, c.Username, c.Password, c.Port,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println("数据库创建失败：", err.Error())
	}

	_, err = db.Exec(`CREATE DATABASE "` + c.Dbname + `"`)
	if err != nil {
		myLog.Log().Infoln(err)
	}
}
