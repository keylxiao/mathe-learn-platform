package controllers

import (
    "github.com/kataras/iris/v12"
    "io"
    "net/http"
    "os"
)

// PostLoadImage 上传用户头像
func PostLoadImage(c iris.Context) {
    userId := c.URLParam("id")
    file, _, err := c.FormFile("image")
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
        c.JSON("上传失败")
        return
    }
    defer file.Close()
    userId = userId + ".jpg"
    out, err := os.OpenFile("./UserImage/"+userId, os.O_WRONLY|os.O_CREATE, 0666)
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
