package controllers

import (
    "github.com/kataras/iris/v12"
    "mathe-learn-platform/models"
    "net/http"
)

// GetHomePageInfo 首页总览信息
func GetHomePageInfo(c iris.Context) {
    result, err := models.GetHomePageInfo()
    if err != nil {
        c.StatusCode(http.StatusInternalServerError)
    } else {
        c.StatusCode(http.StatusOK)
        c.JSON(result)
    }
}

// PutPlatformVisit 访问平台人数增加
func PutPlatformVisit(c iris.Context) {
    models.PutPlatformVisit()
}
