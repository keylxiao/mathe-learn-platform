package controllers

import (
    "github.com/kataras/iris/v12"
    "mathe-learn-platform/models"
    "mathe-learn-platform/utils"
    "net/http"
)

// PostOnloadBlog 博文上传
//func PostOnloadBlog(c iris.Context) {
//    id := utils.GetUUID("blog")
//    var blog models.Blog
//    c.ReadJSON(&blog)
//    blog.BlogId = id
//    file, _, err := c.FormFile("blog")
//    if err != nil {
//        c.StatusCode(http.StatusInternalServerError)
//        c.JSON("上传失败")
//        return
//    }
//    defer file.Close()
//    id = id + ".txt"
//    out, err := os.OpenFile(config.BlogStorageAddress+id, os.O_WRONLY|os.O_CREATE, 0666)
//    if err != nil {
//        c.StatusCode(http.StatusInternalServerError)
//        c.JSON("上传失败")
//        return
//    }
//    defer out.Close()
//    io.Copy(out, file)
//    err = models.PostOnloadBlog(blog)
//    if err != nil {
//        c.StatusCode(http.StatusInternalServerError)
//        c.JSON("上传失败")
//        return
//    }
//    c.StatusCode(http.StatusOK)
//    c.JSON("上传成功")
//}

// PostOnloadBlogInf 博文信息上传
func PostOnloadBlogInf(c iris.Context) {
    id := utils.GetUUID("blog")
    var blog models.Blog
    c.ReadJSON(&blog)
    blog.BlogId = id
    err := models.PostOnloadBlogInf(blog)
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
        c.JSON("上传失败")
        return
    }
    c.StatusCode(http.StatusOK)
    c.JSON(id)
}

// PostOnloadBlogBody 博文主体上传
func PostOnloadBlogBody(c iris.Context) {
    var body models.BlogBody
    c.ReadJSON(&body)
    err := models.PostOnloadBlogBody(body)
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

// GetAllBlogList 获取所有博文目录
func GetAllBlogList(c iris.Context) {
    info, err := models.GetAllBlogList()
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
        return
    } else {
        c.StatusCode(http.StatusOK)
        c.JSON(info)
    }
}

// GetBlogListByKW 按关键字查找博文
func GetBlogListByKW(c iris.Context) {
    kw := c.URLParam("kw")
    result, err := models.GetBlogListByKW(kw)
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
    } else {
        c.StatusCode(http.StatusOK)
        c.JSON(result)
    }
}

//// PutUpdateBlog 修改博文
//func PutUpdateBlog(c iris.Context) {
//    id := c.URLParam("id")
//    rename := c.FormValue("name")
//    reintro := c.FormValue("intro")
//    err := models.PutUpdateBlog(id, rename, reintro)
//    if err != nil {
//        c.StatusCode(http.StatusInternalServerError)
//        c.JSON("修改失败")
//        return
//    }
//    file, _, err := c.FormFile("blog")
//    if err != nil {
//        c.StatusCode(http.StatusInternalServerError)
//        c.JSON("修改失败")
//        return
//    }
//    defer file.Close()
//    id = id + ".txt"
//    out, err := os.OpenFile(config.BlogStorageAddress+id, os.O_WRONLY|os.O_CREATE, 0666)
//    if err != nil {
//        c.StatusCode(http.StatusInternalServerError)
//        c.JSON("修改失败")
//        return
//    }
//    defer out.Close()
//    io.Copy(out, file)
//    c.StatusCode(http.StatusOK)
//    c.JSON("ok")
//}

// PutUpdateBlogInf 修改博文信息
func PutUpdateBlogInf(c iris.Context) {
    id := c.URLParam("id")
    rename := c.FormValue("name")
    reintro := c.FormValue("intro")
    err := models.PutUpdateBlogInf(id, rename, reintro)
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
        c.JSON("修改失败")
        return
    }
    c.StatusCode(http.StatusOK)
    c.JSON("ok")
}

// PutUpdateBlogBody 修改博文主体
func PutUpdateBlogBody(c iris.Context) {
    var body models.BlogBody
    c.ReadJSON(&body)
    err := models.PutUpdateBlogBody(body)
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
        c.JSON("修改失败")
        return
    }
    c.StatusCode(http.StatusOK)
    c.JSON("ok")
}

// PutUpdateBlogState 修改博文状态
func PutUpdateBlogState(c iris.Context) {
    var info models.BlogState
    c.ReadJSON(&info)
    err := models.PutUpdateBlogState(info)
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
    } else {
        c.StatusCode(http.StatusOK)
        c.JSON("ok")
    }
}

// GetViewBlog 查看具体博文
func GetViewBlog(c iris.Context) {
    id := c.URLParam("id")
    body, err := models.GetViewBlog(id)
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
    } else {
        c.StatusCode(http.StatusOK)
        c.JSON(body)
    }
}
