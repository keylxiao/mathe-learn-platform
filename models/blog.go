package models

import (
    "mathe-learn-platform/utils"
    "time"
)

// Blog 博文信息结构体
type Blog struct {
    UserId     string `gorm:"primary_key"` // 用户id
    BlogId     string // 博文id
    BlogName   string // 博文名称
    BriefIntro string // 博文简介
    CreateTime string // 创建时间
    UpdateTime string // 修改时间
}

// PostOnloadBlog 上传用户博文
func PostOnloadBlog(blog Blog) error {
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

// PutUpdateBlog 修改用户博文信息
func PutUpdateBlog(id string, update ...string) error {
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
