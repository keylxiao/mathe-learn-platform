package models

import (
    "context"
    "go.mongodb.org/mongo-driver/bson"
    "gorm.io/gorm"
    "mathe-learn-platform/config"
    "mathe-learn-platform/utils"
    "time"
)

// Bar 帖子信息
type Bar struct {
    Id          string // 帖子id
    UserId      string // 用户id
    Classify    string // 分类
    Name        string // 帖子名称
    Brief       string // 简介
    CreateTime  string // 创建时间
    FloorNumber int    // 总盖楼数
    PageView    int    // 浏览量
    LikesNumber int    // 点赞数
}

// BarFloor 楼层信息
type BarFloor struct {
    Id             string // 楼层id
    BarId          string // 所属帖子id
    UserId         string // 用户id
    CreateTime     string // 创建时间
    LikesNumber    int    // 点赞数
    SonFloorNumber int    // 楼中楼数
}

// SonFloor 楼中楼信息
type SonFloor struct {
    Id          string // 楼中楼id
    BarId       string // 所属帖子id
    BarFloorId  string // 所属楼层id
    Userid      string // 用户id
    ReplyId     string // 回复楼中楼id, 空则为回复所在楼层
    CreateTime  string // 创建时间
    LikesNumber int    // 点赞数
}

// Likes 点赞更新
type Likes struct {
    Id       string // id
    Classify string //更新点赞的目标类型
}

// ReplyBody 楼层及楼中楼回复主体
type ReplyBody struct {
    Id    string // id
    Reply string // 回复信息
}

// BarBody 帖子主体
type BarBody struct {
    Bar           Bar         //帖子信息
    BarFloor      []BarFloor  // 楼层信息
    SonFloor      []SonFloor  // 楼中楼信息
    BarFloorReply []ReplyBody // 楼层回复主体
    SonFloorReply []ReplyBody // 楼中楼回复主体
}

// PostPublishBar 发布新帖
func PostPublishBar(bar Bar) error {
    bar.CreateTime = time.Now().Format("2006-01-02 15:04:05")
    db := utils.DBOpen()
    sqlDB, _ := db.DB()
    defer sqlDB.Close()
    tx := db.Begin()
    err := tx.Create(&bar).Error
    if err != nil {
        tx.Rollback()
        return err
    }
    tx.Commit()
    return err
}

// PostReplyBarInf 楼层信息
func PostReplyBarInf(barfloor BarFloor) error {
    barfloor.CreateTime = time.Now().Format("2006-01-02 15:04:05")
    db := utils.DBOpen()
    sqlDB, _ := db.DB()
    defer sqlDB.Close()
    tx := db.Begin()
    err := tx.Create(&barfloor).Error
    if err != nil {
        tx.Rollback()
        return err
    }
    tx.Commit()
    return err
}

// PostReplyBarBody 楼层主体
func PostReplyBarBody(body ReplyBody) error {
    client := utils.MongoOpen()
    mongo := client.Database(config.MongoDBName)
    collection := mongo.Collection(config.MongoBlogTabName)
    _, err := collection.InsertOne(context.TODO(), body)
    return err
}

// PostReplyBarFloorInf 楼中楼信息
func PostReplyBarFloorInf(sonfloor SonFloor) error {
    sonfloor.CreateTime = time.Now().Format("2006-01-02 15:04:05")
    db := utils.DBOpen()
    sqlDB, _ := db.DB()
    defer sqlDB.Close()
    tx := db.Begin()
    err := tx.Create(&sonfloor).Error
    if err != nil {
        tx.Rollback()
        return err
    }
    tx.Commit()
    return err
}

// PostReplyBarFloorBody 楼中楼主体
func PostReplyBarFloorBody(body ReplyBody) error {
    client := utils.MongoOpen()
    mongo := client.Database(config.MongoDBName)
    collection := mongo.Collection(config.MongoBlogTabName)
    _, err := collection.InsertOne(context.TODO(), body)
    return err
}

// PutBarLikes 点赞统一接口
func PutBarLikes(like Likes) error {
    db := utils.DBOpen()
    sqlDB, _ := db.DB()
    defer sqlDB.Close()
    tx := db.Begin()
    var err error
    switch like.Classify {
    case "bar":
        err = tx.Model(&Bar{}).Where("id", like.Id).Update("likes_number", gorm.Expr("likes_number + ?", 1)).Error
    case "barfloor":
        err = tx.Model(&BarFloor{}).Where("id", like.Id).Update("likes_number", gorm.Expr("likes_number + ?", 1)).Error
    case "sonfloor":
        err = tx.Model(&SonFloor{}).Where("id", like.Id).Update("likes_number", gorm.Expr("likes_number + ?", 1)).Error
    }
    if err != nil {
        tx.Rollback()
        return err
    }
    tx.Commit()
    return err
}

// GetViewBar 查看帖子
func GetViewBar(id string) (BarBody, error) {
    var body BarBody
    // mysql获取信息
    db := utils.DBOpen()
    sqlDB, _ := db.DB()
    defer sqlDB.Close()
    err := db.Where("id = ?", id).Find(&body.Bar).Error
    err = db.Where("bar_id = ?", id).Find(&body.BarFloor).Error
    err = db.Where("bar_id = ?", id).Find(&body.SonFloor).Error
    // mongodb获取主体
    client := utils.MongoOpen()
    mongo := client.Database(config.MongoDBName)
    collection := mongo.Collection(config.MongoBarTabName)
    err = collection.FindOne(context.TODO(), bson.D{{"id", id}}).Decode(&body.BarFloorReply)
    err = collection.FindOne(context.TODO(), bson.D{{"id", id}}).Decode(&body.SonFloorReply)
    return body, err
}
