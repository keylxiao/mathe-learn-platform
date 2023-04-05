package models

import (
    "gorm.io/gorm"
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

// PostReplyBar 回复帖子
func PostReplyBar(barfloor BarFloor) error {
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

// PostReplyBarFloor 楼中楼
func PostReplyBarFloor(sonfloor SonFloor) error {
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
