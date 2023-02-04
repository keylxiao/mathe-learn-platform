package models

import (
    "errors"
    "mathe-learn-platform/utils"
)

type User struct {
    Id         string `gorm:"primary_key"` // 主键
    Account    string // 学号/工号
    Password   string // 密码(sha1加密)
    UserName   string // 用户名
    Telephone  string // 电话号码
    QQNumber   string // QQ号
    Status     int    // 用户权限(0平台管理员 1老师 2学生)
    College    int    // 学院
    Major      string // 专业
    CreateTime string // 创建时间
    UpdateTime string // 修改时间
}

// CheckOnlyOne 检查用户信息是否已被注册
func CheckOnlyOne(info User) string {
    db := utils.DBOpen()
    sqlDB, _ := db.DB()
    defer sqlDB.Close()
    var user User
    db.Where("account = ?", info.Account).Find(&user)
    if user.Account != "" {
        return "该学号已被注册！"
    }
    db.Where("telephone = ?", info.Telephone).Find(&user)
    if user.Account != "" {
        return "该手机号已被注册！"
    }
    db.Where("qq_number = ?", info.QQNumber).Find(&user)
    if user.Account != "" {
        return "该QQ号已被注册！"
    }
    return ""
}

// PostUserRegister 用户注册
func PostUserRegister(info User) error {
    db := utils.DBOpen()
    sqlDB, _ := db.DB()
    defer sqlDB.Close()
    tx := db.Begin()
    err := tx.Create(&info).Error
    if err != nil {
        tx.Rollback()
        return err
    }
    tx.Commit()
    return err
}

// GetUserLogin 用户登录
func GetUserLogin(account, password string) (bool, error) {
    db := utils.DBOpen()
    sqlDB, _ := db.DB()
    defer sqlDB.Close()
    var info User
    err := db.Where("account = ?", account).Find(&info).Error
    if err != nil {
        return false, err
    }
    if info.Password == "" {
        return false, errors.New("none info")
    }
    if info.Password == password {
        return true, err
    } else {
        return false, err
    }
}
