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
    var blog models.Blog
    c.ReadJSON(&blog)
    blog.BlogId = id
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
    err = models.PostOnloadBlog(blog)
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

// PutUpdateBlog 修改博文
func PutUpdateBlog(c iris.Context) {
    id := c.URLParam("id")
    rename := c.FormValue("name")
    reintro := c.FormValue("intro")
    err := models.PutUpdateBlog(id, rename, reintro)
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
        c.JSON("修改失败")
        return
    }
    file, _, err := c.FormFile("blog")
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
        c.JSON("修改失败")
        return
    }
    defer file.Close()
    id = id + ".txt"
    out, err := os.OpenFile(config.BlogStorageAddress+id, os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
        c.JSON("修改失败")
        return
    }
    defer out.Close()
    io.Copy(out, file)
    c.StatusCode(http.StatusOK)
    c.JSON("ok")
}

// GetViewBlog 查看具体博文
func GetViewBlog(c iris.Context) {
    id := c.URLParam("id")
    id = config.BlogStorageAddress + id + ".txt"
    c.StatusCode(http.StatusOK)
    c.SendFile(id, "blog.txt")
}
