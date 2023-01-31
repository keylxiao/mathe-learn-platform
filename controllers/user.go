package controllers

import (
    "github.com/kataras/iris/v12"
    "mime/multipart"
    "net/http"
)

// PostLoadImage 上传用户头像
func PostLoadImage(c iris.Context) {
    userId := c.URLParam("id")
    _, err := c.UploadFormFiles("./UserImage", func(ctx iris.Context, file *multipart.FileHeader) {
        file.Filename = userId
    })
    if err != nil {
        c.StatusCode(http.StatusBadRequest)
        c.JSON("上传失败")
    } else {
        c.StatusCode(http.StatusOK)
        c.JSON("上传成功")
    }
}
