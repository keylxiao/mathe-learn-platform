package models

import (
    "mathe-learn-platform/utils"
)

type HomePageInformation struct {
    TotalVisitor int // 总访问人次
}

// GetHomePageInfo 首页总览信息
func GetHomePageInfo() (HomePageInformation, error) {
    db := utils.DBOpen()
    var result HomePageInformation
    err := db.Find(&result).Error
    sqlDB, _ := db.DB()
    defer sqlDB.Close()
    return result, err
}

// PutPlatformVisit 访问平台人数增加
func PutPlatformVisit() {
    db := utils.DBOpen()
    var info HomePageInformation
    db.Find(&info)
    info.TotalVisitor++
    db.Model(&info).Where("total_visitor = ?", info.TotalVisitor-1).Update("total_visitor", info.TotalVisitor)
    sqlDB, _ := db.DB()
    defer sqlDB.Close()
}
