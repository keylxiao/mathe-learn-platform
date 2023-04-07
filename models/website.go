package models

import (
    "mathe-learn-platform/utils"
)

type OutsideWeb struct {
    Name       string // 网站内容
    Link       string // 网址
    Class      string // 学科门类
    PictureId  string // 图标id
    BriefIntro string // 简介
}

// PostOnloadWeb 添加外部网站
func PostOnloadWeb(web OutsideWeb) error {
    db := utils.DBOpen()
    sqlDB, _ := db.DB()
    defer sqlDB.Close()
    tx := db.Begin()
    err := tx.Create(&web).Error
    if err != nil {
        tx.Rollback()
        return err
    }
    tx.Commit()
    return err
}

// GetWebsiteByKW 按关键字返回外部网站
func GetWebsiteByKW(kw string) ([]OutsideWeb, error) {
    var info []OutsideWeb
    db := utils.DBOpen()
    err := db.Where("name LIKE ?", "%"+kw+"%").Or("class LIKE ?", "%"+kw+"%").Or("brief_intro LIKE ?", "%"+kw+"%").Find(&info).Error
    return info, err
}
