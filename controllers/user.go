package controllers

import (
    "errors"
    "github.com/kataras/iris/v12"
    "io"
    "mathe-learn-platform/config"
    "mathe-learn-platform/models"
    "net/http"
    "os"
)

// PostUserRegister 用户注册
func PostUserRegister(c iris.Context) {
    var info models.User
    c.ReadJSON(&info)
    // 恶意用户权限检测
    if info.Status == 0 {
        c.StatusCode(http.StatusBadRequest)
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

// GetSendCode 用户请求发送验证码
func GetSendCode(c iris.Context) {
    userId := c.URLParam("id")
    err := models.GetSendCode(userId)
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
    } else {
        c.StatusCode(http.StatusOK)
        c.JSON("发送成功")
    }
}

// GetVerifyCode 用户验证码验证
func GetVerifyCode(c iris.Context) {
    userId := c.URLParam("id")
    code := c.URLParam("code")
    isreal, err := models.GetVerifyCode(userId, code)
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
    } else {
        c.StatusCode(http.StatusOK)
        c.JSON(isreal)
    }
}

// GetUserInfo 查看用户信息
func GetUserInfo(c iris.Context) {
    userId := c.URLParam("id")
    info, err := models.GetUserInfo(userId)
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
    } else {
        c.StatusCode(http.StatusOK)
        c.JSON(info)
    }
}

// GetSearchUser 检索用户
func GetSearchUser(c iris.Context) {
    info := c.URLParam("info")
    result, err := models.GetSearchUser(info)
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
    } else {
        c.StatusCode(http.StatusOK)
        c.JSON(result)
    }
}

// PutUpdateUser 修改用户信息
func PutUpdateUser(c iris.Context) {
    var update []models.UpdateUserInfo
    id := c.URLParam("id")
    c.ReadJSON(&update)
    err := models.PutUpdateUser(id, update)
    if err != nil {
        if err == errors.New("exceed authority") {
            c.StatusCode(http.StatusBadRequest)
            c.JSON("exceed authority")
        } else {
            c.StatusCode(http.StatusInternalServerError)
        }
    } else {
        c.StatusCode(http.StatusOK)
        c.JSON("修改成功")
    }
}

// PostOnloadIcon 上传用户头像
func PostOnloadIcon(c iris.Context) {
    userId := c.URLParam("id")
    file, _, err := c.FormFile("image")
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
        c.JSON("上传失败")
        return
    }
    defer file.Close()
    userId = userId + ".jpg"
    out, err := os.OpenFile(config.IconStorageAddress+userId, os.O_WRONLY|os.O_CREATE, 0666)
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

// GetUserIcon 获取用户头像
func GetUserIcon(c iris.Context) {
    userId := c.URLParam("id")
    userId = config.IconStorageAddress + userId + ".jpg"
    c.StatusCode(http.StatusOK)
    c.SendFile(userId, "icon.jpg")
}
