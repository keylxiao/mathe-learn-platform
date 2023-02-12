package controllers

import (
    "github.com/kataras/iris/v12"
    "io"
    "mathe-learn-platform/config"
    "mathe-learn-platform/models"
    "mathe-learn-platform/utils"
    "net/http"
    "os"
)

// PostOnloadBlog 博文上传
func PostOnloadBlog(c iris.Context) {
    id := utils.GetUUID("blog")
    userId := c.URLParam("id")
    file, _, err := c.FormFile("blog")
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
        c.JSON("上传失败")
        return
    }
    defer file.Close()
    id = id + ".txt"
    out, err := os.OpenFile(config.BlogStorageAddress+id, os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
        c.JSON("上传失败")
        return
    }
    defer out.Close()
    io.Copy(out, file)
    err = models.PostOnloadBlog(userId, id)
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
        c.JSON("上传失败")
        return
    }
    c.StatusCode(http.StatusOK)
    c.JSON("上传成功")
}

// GetUserBlogList 获取用户博文目录
func GetUserBlogList(c iris.Context) {
    id := c.URLParam("id")
    info, err := models.GetUserBlogList(id)
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
        return
    } else {
        c.StatusCode(http.StatusOK)
        c.JSON(info)
    }
}
