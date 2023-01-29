package models

import (
    "gorm.io/gorm"
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
    db.Model(&HomePageInformation{}).Update("total_visitor", gorm.Expr("total_visitor + ?", 1))
    sqlDB, _ := db.DB()
    defer sqlDB.Close()
}
