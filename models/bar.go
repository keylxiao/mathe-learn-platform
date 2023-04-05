package models

// Bar 帖子信息
type Bar struct {
    Id          string // 帖子id
    UserId      string // 用户id
    Classify    string // 分类
    Name        string // 帖子名称
    Brief       string // 简介
    CreateTime  string // 创建时间
    FloorNumber int    // 总盖楼数
    PageView    int    // 浏览量
    LikesNumber int    // 点赞数
}

// BarFloor 楼层信息
type BarFloor struct {
    Id             string // 楼层id
    BarId          string // 所属帖子id
    UserId         string // 用户id
    CreateTime     string // 创建时间
    LikesNumber    int    // 点赞数
    SonFloorNumber int    // 楼中楼数
}

// SonFloor 楼中楼信息
type SonFloor struct {
    Id          string // 楼中楼id
    BarFloorId  string // 所属楼层id
    Userid      string // 用户id
    ReplyId     string // 回复楼中楼id, 空则为回复所在楼层
    CreateTime  string // 创建时间
    LikesNumber int    // 点赞数
}
