package router

import (
    "github.com/kataras/iris/v12"
    "mathe-learn-platform/controllers"
)

// HomePageRouters 主页路由
func HomePageRouters(home iris.Party) {
    // 首页总览信息
    home.Get("/HomePageInfo", controllers.GetHomePageInfo)
    // 访问人数增加
    home.Put("/PlatformVisit", controllers.PutPlatformVisit)
}

// UserRouter 用户路由
func UserRouter(user iris.Party) {
    // 用户注册
    user.Post("/UserRegister", controllers.PostUserRegister)
    // 用户登录
    user.Get("/UserLogin", controllers.GetUserLogin)
    // 查看用户信息
    user.Get("/UserInfo", controllers.GetUserInfo)
    // 获取用户头像
    user.Get("/UserIcon", controllers.GetUserIcon)
    // 检索用户
    user.Get("/SearchUser", controllers.GetSearchUser)
    // 修改用户信息
    user.Put("/UpdateUser", controllers.PutUpdateUser)
    // 用户头像上传
    user.Post("/OnloadIcon", controllers.PostOnloadIcon)
    // 用户请求发送验证码
    user.Get("/SendCode", controllers.GetSendCode)
    // 用户验证码验证
    user.Get("/VerifyCode", controllers.GetVerifyCode)
}

// WebsiteRouter 网站路由
func WebsiteRouter(web iris.Party) {
    // 按关键字返回外部网站
    web.Get("/OutsideWeb", controllers.GetWebsiteByKW)
}

// BlogRouter 博文路由
func BlogRouter(blog iris.Party) {
    // 博文上传
    blog.Post("/OnloadBlog", controllers.PostOnloadBlog)
    // 博文修改
    blog.Put("/UpdateBlog", controllers.PutUpdateBlog)
    // 博文权限修改

    // 查看用户博文目录
    blog.Get("/UserBlogList", controllers.GetUserBlogList)
    // 查看具体博文
}

// PostBarRouter 贴吧路由
func PostBarRouter(bar iris.Party) {
    // 发布新帖

    // 回复帖子

    // 消息盒子

    // 删除帖子
}
