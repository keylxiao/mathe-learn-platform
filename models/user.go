package models

type User struct {
    Id         int    `gorm:"primary_key;auto_increment"` // 自增主键
    Account    string // 登录账号
    Password   string // 密码(sha1加密)
    UserName   string // 用户名
    TelePhone  string // 电话号码
    Status     int    // 用户权限(0平台管理员 1老师 2学生)
    CreateTime string // 创建时间
    UpdateTime string // 修改时间
    IsDelete   int    // 逻辑删除(0未删 1删除)
}
