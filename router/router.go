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
    // 添加外部网站
    web.Post("/OnloadWeb", controllers.PostOnloadWeb)
    // 添加外部网站图片
    web.Post("/OnloadWebPicture", controllers.PostOnloadPic)
    // 按关键字返回外部网站
    web.Get("/OutsideWeb", controllers.GetWebsiteByKW)
    // 获取外部网站图片
    web.Get("/WebPicture", controllers.GetWebPicture)
}

// BlogRouter 博文路由
func BlogRouter(blog iris.Party) {
    // 博文信息上传
    blog.Post("/OnloadBlogInf", controllers.PostOnloadBlogInf)
    // 博文主体上传
    blog.Post("/OnloadBlogBody", controllers.PostOnloadBlogBody)
    // 博文信息修改
    blog.Put("/UpdateBlogInf", controllers.PutUpdateBlogInf)
    // 博文主体修改
    blog.Put("/UpdateBlogBody", controllers.PutUpdateBlogBody)
    // 博文权限修改
    blog.Put("/UpdateBlogState", controllers.PutUpdateBlogState)
    // 查看用户博文目录
    blog.Get("/UserBlogList", controllers.GetUserBlogList)
    // 获取所有博文目录
    blog.Get("/AllBlogList", controllers.GetAllBlogList)
    // 获取博文信息
    blog.Get("/BlogInf", controllers.GetBlogInf)
    // 按关键字获取博文目录
    blog.Get("/BlogListByKW", controllers.GetBlogListByKW)
    // 查看具体博文
    blog.Get("/ViewBlog", controllers.GetViewBlog)
}

// PostBarRouter 贴吧路由
func PostBarRouter(bar iris.Party) {
    // 发布新帖
    bar.Post("/PublishBar", controllers.PostPublishBar)
    // 楼层信息
    bar.Post("/ReplyBarInf", controllers.PostReplyBarInf)
    // 楼层主体
    bar.Post("/ReplyBarBody", controllers.PostReplyBarBody)
    // 楼中楼信息
    bar.Post("/ReplyBarFloorInf", controllers.PostReplyBarFloorInf)
    // 楼中楼主体
    bar.Post("/ReplyBarFloorBody", controllers.PostReplyBarFloorBody)
    // 点赞统一接口
    bar.Put("/BarLikes", controllers.PutBarLikes)
    // 消息盒子

    // 查看帖子
    bar.Get("/ViewBar", controllers.GetViewBar)
    // 删除帖子
}
