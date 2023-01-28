package router

import "github.com/kataras/iris/v12"

// HomePageRouters 主页路由
func HomePageRouters(home iris.Party) {
    // 首页总览信息
}

// UserRouter 用户路由
func UserRouter(user iris.Party) {
    // 用户注册

    // 用户登录

    // 查看用户信息

    // 修改用户信息
}

// WebsiteRouter 网站路由
func WebsiteRouter(web iris.Party) {
    // 外部网站推荐

}

// BlogRouter 博文路由
func BlogRouter(blog iris.Party) {
    // 博文上传

    // 博文修改

    // 博文权限修改

    // 查看用户博文目录

    // 查看具体博文
}

// PostBarRouter 贴吧路由
func PostBarRouter(bar iris.Party) {
    // 发布新帖

    // 回复帖子

    // 消息盒子

    // 删除帖子
}
