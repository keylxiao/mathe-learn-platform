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

// PostReplyBarInf 楼层信息
func PostReplyBarInf(c iris.Context) {
    id := utils.GetUUID("barfloor")
    var barfloor models.BarFloor
    c.ReadJSON(&barfloor)
    barfloor.Id = id
    err := models.PostReplyBarInf(barfloor)
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
        c.JSON("上传失败")
        return
    }
    c.StatusCode(http.StatusOK)
    c.JSON(id)
}

// PostReplyBarBody 楼层主体
func PostReplyBarBody(c iris.Context) {
    var body models.ReplyBody
    c.ReadJSON(&body)
    err := models.PostReplyBarBody(body)
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
        c.JSON("上传失败")
        return
    }
    c.StatusCode(http.StatusOK)
    c.JSON("上传成功")
}

// PostReplyBarFloorInf 楼中楼信息
func PostReplyBarFloorInf(c iris.Context) {
    id := utils.GetUUID("sonfloor")
    var son models.SonFloor
    c.ReadJSON(&son)
    son.Id = id
    err := models.PostReplyBarFloorInf(son)
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
        c.JSON("上传失败")
        return
    }
    c.StatusCode(http.StatusOK)
    c.JSON(id)
}

// PostReplyBarFloorBody 楼层主体
func PostReplyBarFloorBody(c iris.Context) {
    var body models.ReplyBody
    c.ReadJSON(&body)
    err := models.PostReplyBarFloorBody(body)
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
        c.JSON("上传失败")
        return
    }
    c.StatusCode(http.StatusOK)
    c.JSON("上传成功")
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

// GetViewBar 查看帖子
func GetViewBar(c iris.Context) {
    id := c.URLParam("id")
    body, err := models.GetViewBar(id)
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
    } else {
        c.StatusCode(http.StatusOK)
        c.JSON(body)
    }
}
