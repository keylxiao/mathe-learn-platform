package models

import (
    "errors"
    "math/rand"
    "mathe-learn-platform/utils"
    "strconv"
    "time"
)

// User 用户信息结构体
type User struct {
    Id         string `gorm:"primary_key"` // 主键
    Account    string // 学号/工号
    Password   string // 密码(sha1加密)
    UserName   string // 用户名
    Telephone  string // 电话号码
    Mailbox    string // 邮箱地址
    Status     int    // 用户权限(0平台管理员 1老师 2学生)
    College    int    // 学院
    Major      string // 专业
    CreateTime string // 创建时间
    UpdateTime string // 修改时间
}

// UpdateUserInfo 用户修改信息结构体
type UpdateUserInfo struct {
    UpdateField string      // 更新字段
    NewInfo     interface{} // 字段更新值
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
    db.Where("mailbox = ?", info.Mailbox).Find(&user)
    if user.Account != "" {
        return "该邮箱已被注册！"
    }
    return ""
}

// PostUserRegister 用户注册
func PostUserRegister(info User) error {
    info.Id = utils.GetUUID("user")
    info.CreateTime = time.Now().Format("2006-01-02 15:04:05")
    info.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
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
func GetUserLogin(account, password string) (bool, error, bool) {
    db := utils.DBOpen()
    sqlDB, _ := db.DB()
    defer sqlDB.Close()
    var info User
    err := db.Where("account = ?", account).Find(&info).Error
    if err != nil {
        return false, err, false
    }
    if info.Password == "" {
        return false, errors.New("none info"), true
    }
    if info.Password == password {
        return true, err, true
    } else {
        return false, err, true
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
    err = utils.SendEmail(user.Mailbox, code, user.UserName)
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

// GetUserInfo 查看用户信息
func GetUserInfo(id string) (User, error) {
    db := utils.DBOpen()
    sqlDB, _ := db.DB()
    defer sqlDB.Close()
    var info User
    err := db.Where("id = ?", id).Find(&info).Error
    if err != nil {
        return info, err
    }
    info.Password = ""
    return info, nil
}

// GetSearchUser 检索用户
func GetSearchUser(info string) (User, error) {
    db := utils.DBOpen()
    sqlDB, _ := db.DB()
    defer sqlDB.Close()
    var user User
    err := db.Where("user_name LIKE ?", info).Or("college = ?", info).Or("major LIKE ?", info).Find(&user).Error
    return user, err
}

// PutUpdateUser 修改用户信息
func PutUpdateUser(id string, update []UpdateUserInfo) (err error) {
    db := utils.DBOpen()
    sqlDB, _ := db.DB()
    defer sqlDB.Close()
    tx := db.Begin()
    var user User
    for i := range update {
        if update[i].UpdateField == "Id" || update[i].UpdateField == "Status" || update[i].UpdateField == "CreateTime" {
            return errors.New("exceed authority")
        }
        if update[i].UpdateField == "College" {
            update[i].NewInfo = update[i].NewInfo.(int)
        } else {
            update[i].NewInfo = update[i].NewInfo.(string)
        }
        err = tx.Model(&user).Where("id = ?", id).Update(update[i].UpdateField, update[i].NewInfo).Error
        if err != nil {
            tx.Rollback()
            return
        }
    }
    tx.Commit()
    return
}
