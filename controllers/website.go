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

// PostOnloadWeb 添加外部网站
func PostOnloadWeb(c iris.Context) {
    id := utils.GetUUID("web")
    var web models.OutsideWeb
    c.ReadJSON(&web)
    web.PictureId = id
    err := models.PostOnloadWeb(web)
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
    } else {
        c.StatusCode(http.StatusOK)
        c.JSON(id)
    }
}

// PostOnloadPic 上传网站图片
func PostOnloadPic(c iris.Context) {
    userId := c.FormValue("id")
    file, _, err := c.FormFile("image")
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
        c.JSON("上传失败")
        return
    }
    defer file.Close()
    userId = userId + ".jpg"
    out, err := os.OpenFile(config.PicStorageAddress+userId, os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
        c.JSON("上传失败")
        return
    }
    defer out.Close()
    io.Copy(out, file)
    c.StatusCode(http.StatusOK)
    c.JSON("上传成功")
}

// GetWebsiteByKW 按关键字返回外部网站
func GetWebsiteByKW(c iris.Context) {
    kw := c.URLParam("kw")
    result, err := models.GetWebsiteByKW(kw)
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
    } else {
        c.StatusCode(http.StatusOK)
        c.JSON(result)
    }
}

// GetWebPicture 获取外部网站图片
func GetWebPicture(c iris.Context) {
    userId := c.URLParam("id")
    userId = config.IconStorageAddress + userId + ".jpg"
    c.StatusCode(http.StatusOK)
    c.SendFile(userId, "icon.jpg")
}
