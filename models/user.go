package models

import (
    "errors"
    "math/rand"
    "mathe-learn-platform/utils"
    "strconv"
    "time"
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

// GetSendCode 用户请求发送验证码
func GetSendCode(id string) error {
    db := utils.RedisOpen()
    defer db.Close()
    sql := utils.DBOpen()
    sqlDB, _ := sql.DB()
    defer sqlDB.Close()
    var user User
    err := sql.Where("id = ?", id).Find(&user).Error
    if err != nil {
        return err
    }
    code := strconv.Itoa(rand.Int() % 10000)
    err = db.Set(id, code, 5*time.Minute).Err()
    if err != nil {
        return err
    }
    err = utils.SendEmail(user.QQNumber, code, user.UserName)
    if err != nil {
        return err
    }
    return nil
}

// GetVerifyCode 用户验证码验证
func GetVerifyCode(id, code string) (bool, error) {
    db := utils.RedisOpen()
    defer db.Close()
    realcode, err := db.Get(id).Result()
    if err != nil {
        return false, err
    }
    if realcode == code {
        return true, err
    } else {
        return false, err
    }
}
