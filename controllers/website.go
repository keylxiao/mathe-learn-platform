package controllers

import (
    "github.com/kataras/iris/v12"
    "mathe-learn-platform/models"
    "net/http"
)

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
