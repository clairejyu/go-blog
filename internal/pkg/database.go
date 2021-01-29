package pkg

import (
	"fmt"
	"time"

	"github.com/clairejyu/go-blog/internal/pkg/setting"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

// InitDB Opening a database and save the reference to `Database` struct.
func InitDB() *gorm.DB {
	dsn := "host=" + setting.DatabaseSetting.Host + " user=" + setting.DatabaseSetting.User + " password=" +
		setting.DatabaseSetting.Password + " dbname=" + setting.DatabaseSetting.Name + " port=" + setting.DatabaseSetting.Port +
		" sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		fmt.Println("db err: ", err)
	}
	sqlDB, err := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifeti	1qme 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	//db.LogMode(true)
	DB = db
	return DB
}

// GetDB Using this function to get a connection, you can create your connection pool here.
// func GetDB() *gorm.DB {
// 	return DB
// }
