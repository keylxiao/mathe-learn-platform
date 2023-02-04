package controllers

import (
    "errors"
    "github.com/kataras/iris/v12"
    "io"
    "mathe-learn-platform/models"
    "mathe-learn-platform/utils"
    "net/http"
    "os"
    "time"
)

// PostUserRegister 用户注册
func PostUserRegister(c iris.Context) {
    var info models.User
    c.ReadJSON(&info)
    // 恶意用户权限检测
    if info.Status == 0 {
        c.StatusCode(http.StatusOK)
        c.JSON("用户恶意修改权限")
        return
    }
    // 用户注册信息查重
    only := models.CheckOnlyOne(info)
    if only != "" {
        c.StatusCode(http.StatusOK)
        c.JSON(only)
        return
    }
    info.Id = utils.GetUUID("user")
    info.CreateTime = time.Now().Format("2006-01-02 15:04:05")
    info.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
    err := models.PostUserRegister(info)
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
        c.JSON("注册失败, 请稍后重试!")
    } else {
        c.StatusCode(http.StatusOK)
        c.JSON("注册成功, 欢迎您!")
    }
}

// GetUserLogin 用户登录
func GetUserLogin(c iris.Context) {
    var info models.User
    c.ReadJSON(&info)
    isReal, err := models.GetUserLogin(info.Account, info.Password)
    if err != nil {
        if err == errors.New("none info") {
            c.StatusCode(http.StatusOK)
            c.JSON("none")
        } else {
            c.StatusCode(http.StatusInternalServerError)
        }
        return
    }
    if isReal {
        c.StatusCode(http.StatusOK)
        c.JSON("true")
    } else {
        c.StatusCode(http.StatusOK)
        c.JSON("false")
    }
}

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
