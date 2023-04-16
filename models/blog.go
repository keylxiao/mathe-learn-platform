package models

import (
    "context"
    "go.mongodb.org/mongo-driver/bson"
    "mathe-learn-platform/config"
    "mathe-learn-platform/utils"
    "time"
)

// Blog 博文信息结构体
type Blog struct {
    UserId     string `gorm:"primary_key"` // 用户id
    BlogId     string // 博文id
    BlogName   string // 博文名称
    BriefIntro string // 博文简介
    State      int    // 博文状态(初始0 仅自己可见1 软删除2)
    CreateTime string // 创建时间
    UpdateTime string // 修改时间
}

// BlogBody 博文主体
type BlogBody struct {
    BlogId string // 博文id
    Body   string // 博文主体

}

// BlogState 修改博文状态结构体
type BlogState struct {
    BlogId   string //博文id
    NewState int    // 新状态
}

// PostOnloadBlogInf 上传用户博文信息
func PostOnloadBlogInf(blog Blog) error {
    // 信息部分存储
    blog.CreateTime = time.Now().Format("2006-01-02 15:04:05")
    blog.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
    db := utils.DBOpen()
    sqlDB, _ := db.DB()
    defer sqlDB.Close()
    tx := db.Begin()
    err := tx.Create(&blog).Error
    if err != nil {
        tx.Rollback()
        return err
    }
    tx.Commit()
    return err
}

// PostOnloadBlogBody 上传用户博文主体
func PostOnloadBlogBody(body BlogBody) error {
    // 博文主体部分存储
    client := utils.MongoOpen()
    mongo := client.Database(config.MongoDBName)
    collection := mongo.Collection(config.MongoBlogTabName)
    _, err := collection.InsertOne(context.TODO(), body)
    if err != nil {
        db := utils.DBOpen()
        sqlDB, _ := db.DB()
        defer sqlDB.Close()
        db.Where("blog_id = ?", body.BlogId).Delete(&Blog{})
    }
    return err
}

// GetUserBlogList 获取用户博文目录
func GetUserBlogList(id string) ([]Blog, error) {
    db := utils.DBOpen()
    sqlDB, _ := db.DB()
    defer sqlDB.Close()
    var blog []Blog
    err := db.Where("user_id = ?", id).Where("state = ?", 0).Find(&blog).Error
    return blog, err
}

// PutUpdateBlogInf 修改用户博文信息
func PutUpdateBlogInf(id string, update ...string) error {
    var blog Blog
    db := utils.DBOpen()
    sqlDB, _ := db.DB()
    defer sqlDB.Close()
    tx := db.Begin()
    err := db.Model(&blog).Where("blog_id = ?", id).Update("create_time", time.Now().Format("2006-01-02 15:04:05")).Error
    if err != nil {
        tx.Rollback()
        return err
    }
    if update[0] != "" {
        err = db.Model(&blog).Where("blog_id = ?", id).Update("blog_name", update[0]).Error
        if err != nil {
            tx.Rollback()
            return err
        }
    }
    if update[1] != "" {
        err = db.Model(&blog).Where("blog_id = ?", id).Update("brief_intro", update[1]).Error
        if err != nil {
            tx.Rollback()
            return err
        }
    }
    tx.Commit()
    return err
}

// PutUpdateBlogBody 修改用户博文主体
func PutUpdateBlogBody(body BlogBody) error {
    client := utils.MongoOpen()
    mongo := client.Database(config.MongoDBName)
    collection := mongo.Collection(config.MongoBlogTabName)
    _, err := collection.UpdateOne(context.TODO(), bson.D{{"blogid", body.BlogId}}, bson.D{{"$set", bson.D{{"body", body.Body}}}})
    return err
}

// PutUpdateBlogState 修改博文状态
func PutUpdateBlogState(info BlogState) error {
    var blog Blog
    db := utils.DBOpen()
    sqlDB, _ := db.DB()
    defer sqlDB.Close()
    tx := db.Begin()
    err := db.Model(&blog).Where("blog_id = ?", info.BlogId).Update("state", info.NewState).Error
    if err != nil {
        tx.Rollback()
        return err
    }
    tx.Commit()
    return err
}

// GetViewBlog 查看具体博文
func GetViewBlog(id string) (string, error) {
    var body BlogBody
    client := utils.MongoOpen()
    mongo := client.Database(config.MongoDBName)
    collection := mongo.Collection(config.MongoBlogTabName)
    err := collection.FindOne(context.TODO(), bson.D{{"blogid", id}}).Decode(&body)
    return body.Body, err
}
