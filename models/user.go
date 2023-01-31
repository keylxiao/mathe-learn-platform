package models

type User struct {
    Id         string `gorm:"primary_key"` // 主键
    Account    string // 学号/工号
    Password   string // 密码(sha1加密)
    UserName   string // 用户名
    Telephone  string // 电话号码
    QQNumebr   string // QQ号
    Status     int    // 用户权限(0平台管理员 1老师 2学生)
    College    int    // 学院
    Major      string // 专业
    CreateTime string // 创建时间
    UpdateTime string // 修改时间
}
