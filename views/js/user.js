function blog_structure(id) {
    $("#Website").empty();
    var body = document.getElementById('Website')
    // 获取文章信息
    $.get('http://127.0.0.1:8001/Blog/BlogInf?id=' + id, function (data) {
        var inf = eval(data)

        var dad_div = document.createElement('div')
        dad_div.className = "sidebar-right-blogdiv"
        body.appendChild(dad_div)

        var name = document.createElement('h2')
        name.innerHTML = inf["BlogName"]
        name.align = "center"
        dad_div.appendChild(name)

        var username = document.createElement('p')
        username.innerHTML = "作者:" + inf["UserId"]
        username.align = "center"
        dad_div.appendChild(username)

        var time = document.createElement('p')
        time.innerHTML = "创建时间:" + inf["CreateTime"]
        time.align = "center"
        dad_div.appendChild(time)

        // 获取文章主体
        $.get('http://127.0.0.1:8001/Blog/ViewBlog?id=' + id, function (data) {
            var page = document.createElement('p')
            page.innerHTML = data
            page.align = "center"
            dad_div.appendChild(page)
        })

        var tr5 = document.createElement('div')
        dad_div.appendChild(tr5)
    })
}
function structure(obj) {
    var body = document.getElementById('Website')

    var dad_div = document.createElement('div')
    body.appendChild(dad_div)
    dad_div.className = "single-member effect-3"

    dad_div.onclick = function () {
        blog_structure(obj["BlogId"])
    }

    var img_div = document.createElement('div')
    dad_div.appendChild(img_div)
    img_div.className = "member-image"

    var img = document.createElement('img')
    img_div.appendChild(img)
    img.src = "http://127.0.0.1:8001/User/UserIcon?id=" + obj["UserId"]

    var inf_div = document.createElement('div')
    dad_div.appendChild(inf_div)
    inf_div.className = "member-info"

    var h3 = document.createElement('h3')
    inf_div.appendChild(h3)
    h3.innerHTML = obj["BlogName"]

    var p = document.createElement('p')
    inf_div.appendChild(p)
    p.innerHTML = obj["BriefIntro"]
}
function web_structure(obj) {
    var body = document.getElementById('Website')

    var dad_div = document.createElement('div')
    body.appendChild(dad_div)
    dad_div.className = "single-member effect-3"

    dad_div.onclick = function () {
        window.location.href = obj["Link"]
    }

    var img_div = document.createElement('div')
    dad_div.appendChild(img_div)
    img_div.className = "member-image"

    var img = document.createElement('img')
    img_div.appendChild(img)
    img.src = "http://127.0.0.1:8001/Website/WebPicture?id=" + obj["PictureId"]

    var inf_div = document.createElement('div')
    dad_div.appendChild(inf_div)
    inf_div.className = "member-info"

    var h3 = document.createElement('h3')
    inf_div.appendChild(h3)
    h3.innerHTML = obj["Name"]

    var p = document.createElement('p')
    inf_div.appendChild(p)
    p.innerHTML = obj["BriefIntro"]
}
function myblog_structure(obj) {
    var body = document.getElementById("Website")
    var dad_div = document.createElement('div')
    body.appendChild(dad_div)
    dad_div.className = "single-member effect-3"

    dad_div.onclick = function () {
        window.location.href = obj["Link"]
    }

    var img_div = document.createElement('div')
    dad_div.appendChild(img_div)
    img_div.className = "member-image"

    var img = document.createElement('img')
    img_div.appendChild(img)
    img.src = "http://127.0.0.1:8001/User/UserIcon?id=" + obj["UserId"]

    var inf_div = document.createElement('div')
    dad_div.appendChild(inf_div)
    inf_div.className = "member-info"

    var h3 = document.createElement('h3')
    inf_div.appendChild(h3)
    h3.innerHTML = obj["BlogName"]

    var p = document.createElement('p')
    inf_div.appendChild(p)
    if (obj["State"] == 0) {
        p.innerHTML = "待审核"
    } else if (obj["State"] == 1) {
        p.innerHTML = "审核通过"
    } else {
        p.innerHTML = "审核未通过"
    }
}