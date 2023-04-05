package controllers

import (
    "github.com/kataras/iris/v12"
    "mathe-learn-platform/models"
    "mathe-learn-platform/utils"
    "net/http"
)

// PostPublishBar 发布新帖
func PostPublishBar(c iris.Context) {
    id := utils.GetUUID("bar")
    var bar models.Bar
    c.ReadJSON(&bar)
    bar.Id = id
    err := models.PostPublishBar(bar)
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
        c.JSON("上传失败")
        return
    }
    c.StatusCode(http.StatusOK)
}

// PostReplyBar 回复帖子
func PostReplyBar(c iris.Context) {
    id := utils.GetUUID("barfloor")
    var barfloor models.BarFloor
    c.ReadJSON(&barfloor)
    barfloor.Id = id
    err := models.PostReplyBar(barfloor)
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
        c.JSON("上传失败")
        return
    }
    c.StatusCode(http.StatusOK)
}

// PostReplyBarFloor 楼中楼
func PostReplyBarFloor(c iris.Context) {
    id := utils.GetUUID("sonfloor")
    var son models.SonFloor
    c.ReadJSON(&son)
    son.Id = id
    err := models.PostReplyBarFloor(son)
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
        c.JSON("上传失败")
        return
    }
    c.StatusCode(http.StatusOK)
}

// PutBarLikes 点赞统一接口
func PutBarLikes(c iris.Context) {
    var like models.Likes
    c.ReadJSON(&like)
    err := models.PutBarLikes(like)
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
        c.JSON("点赞失败")
        return
    }
    c.StatusCode(http.StatusOK)
}
