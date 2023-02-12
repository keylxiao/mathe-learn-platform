package models

import (
    "mathe-learn-platform/utils"
    "time"
)

// Blog 博文信息结构体
type Blog struct {
    UserId     string `gorm:"primary_key"` // 用户id
    BlogId     string // 博文id
    CreateTime string // 创建时间
    UpdateTime string // 修改时间
}

// PostOnloadBlog 上传用户博文
func PostOnloadBlog(userid, blogid string) error {
    var blog Blog
    blog.UserId = userid
    blog.BlogId = blogid
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

// GetUserBlogList 获取用户博文目录
func GetUserBlogList(id string) ([]Blog, error) {
    db := utils.DBOpen()
    sqlDB, _ := db.DB()
    defer sqlDB.Close()
    var blog []Blog
    err := db.Where("user_id = ?", id).Find(&blog).Error
    return blog, err
}
