package main

import (
    "github.com/kataras/iris/v12"
    "github.com/kataras/iris/v12/middleware/recover"
    "mathe-learn-platform/config"
    "mathe-learn-platform/router"
)

func Cors(ctx iris.Context) {
    ctx.Header("Access-Control-Allow-Origin", "*")
    if ctx.Method() == "OPTIONS" {
        ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
        ctx.Header("Access-Control-Allow-Headers", "Content-Type,Accept,Authorization")
        ctx.StatusCode(204)
        return
    }
    ctx.Next()
}
func main() {
    app := iris.New()
    app.Use(recover.New())
    app.UseGlobal(Cors)
    common := app.Party("/")
    {
        common.Options("*", func(ctx iris.Context) {
            ctx.Next()
        })
    }
    // 定义接口路由
    app.PartyFunc("/HomePage", router.HomePageRouters)
    app.PartyFunc("/User", router.UserRouter)
    app.PartyFunc("/Website", router.WebsiteRouter)
    app.PartyFunc("/Blog", router.BlogRouter)
    app.PartyFunc("/PostBar", router.PostBarRouter)
    // 开始监听
    app.Run(iris.Addr(config.BackendAddress+config.Port), iris.WithPostMaxMemory(config.FileSizeLimit))
}
