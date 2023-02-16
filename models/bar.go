package models

// Bar 帖子信息
type Bar struct {
    Id          string // 帖子id
    UserId      string // 用户id
    Name        string // 帖子名称
    Brief       string // 简介
    CreateTime  string // 创建时间
    FloorNumber int    // 总盖楼数
    PageView    int    // 浏览量
    LikesNumber int    // 点赞数
}
