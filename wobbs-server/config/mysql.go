package config

import (
	"fmt"
	"log"
	"os"
	"time"
	"wobbs-server/pkg/snowflake"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"wobbs-server/model"
)

var DB *gorm.DB

func InitDB(conf *MySQLConfig) *gorm.DB {
	username := conf.Username
	password := conf.Password
	host := conf.Host
	port := conf.Port
	database := conf.Database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		database,
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 彩色打印
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_",
			SingularTable: true, // 使用单数
		},
	})

	if err != nil {
		zap.L().Error("failed to connect database ,err:" + err.Error())
		panic(err)
	}
	// 连接池
	sqlDB, err := db.DB()
	if err != nil {
		zap.L().Error("connect db server failed, err:" + err.Error())
		panic(err)
	}
	sqlDB.SetMaxIdleConns(10)           // 设置连接池中空闲连接的最大数量
	sqlDB.SetMaxOpenConns(100)          //设置打开数据库连接的最大数量
	sqlDB.SetConnMaxLifetime(time.Hour) //设置了连接可复用的最大时间
	//数据库迁移
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Category{})
	db.Create(&model.User{BaseModel: model.BaseModel{ID: 1},
		UserID:   snowflake.GenerateID(),
		Username: "admin", Password: "admin",
		Gender: "male",
		Age:    11})
	db.Create(&model.User{BaseModel: model.BaseModel{ID: 2},
		UserID:   snowflake.GenerateID(),
		Username: "novo", Password: "novo",
		Gender: "male",
		Age:    11})
	db.Create(&model.Category{BaseModel: model.BaseModel{ID: 1}, Name: "Go", Description: "Go社区"})
	db.Create(&model.Category{BaseModel: model.BaseModel{ID: 2}, Name: "Java", Description: "Java社区"})
	db.Create(&model.Category{BaseModel: model.BaseModel{ID: 3}, Name: "LeetCode", Description: "算法社区"})
	db.Create(&model.Category{BaseModel: model.BaseModel{ID: 4}, Name: "Acwing", Description: "算法社区"})

	db.AutoMigrate(&model.Post{})
	//db.AutoMigrate(&model.Resource{})
	//db.AutoMigrate(&model.Liked{})       //点赞表
	//db.AutoMigrate(&model.Collect{})     //收藏表
	//db.AutoMigrate(&model.Collection{})  //收藏夹表
	//db.AutoMigrate(&model.Follow{})      //关注表
	//db.AutoMigrate(&model.Comment{})     //评论表
	//db.AutoMigrate(&model.Announce{})    //公告表
	//db.AutoMigrate(&model.UserMessage{}) //私信表
	//db.AutoMigrate(&model.Danmaku{})     //弹幕表
	//db.AutoMigrate(&model.Carousel{})    //轮播图表
	zap.L().Info("database init success")
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
