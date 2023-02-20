package utils

import (
    "context"
    "fmt"
    "github.com/go-redis/redis"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "mathe-learn-platform/config"
)

// BackDBInstance 数据库配置
func BackDBInstance() string {
    command := fmt.Sprintf("%s:%s@(%s)/%s?charset=%s&parseTime=%s&loc=%s", config.DBUserName, config.DBPassword, config.DBRemote, config.DBName, config.DBCharset, config.DBParseTime, config.DBLoc)
    return command
}

// DBOpen 打开Mysql
func DBOpen() *gorm.DB {
    db, _ := gorm.Open(mysql.Open(BackDBInstance()), &gorm.Config{})
    return db
}

// RedisOpen 打开Redis
func RedisOpen() *redis.Client {
    rdb := redis.NewClient(&redis.Options{
        Addr:     config.RedisRemote,
        Password: config.RedisPassword,
        DB:       config.RedisName,
    })
    return rdb
}

// MongoOpen 打开MongoDB
func MongoOpen() *mongo.Client {
    client, _ := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.MongoAddress))
    return client
}
