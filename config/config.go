package config

// 基本参数配置
const (
    // BackendAddress 后端地址
    BackendAddress string = "localhost"
    // Port 端口
    Port string = ":8001"
    // FileSizeLimit 上传文件最大限制
    FileSizeLimit int64 = 1024
    // MLPEmailAddress MLP邮箱总地址
    MLPEmailAddress string = "1285843834@qq.com"
    // EmailCode MLP邮箱授权码
    EmailCode string = "jzruajinzbyzhghj"
    // IconStorageAddress 用户头像存储地址
    IconStorageAddress string = "./UserImage/"
    // BlogStorageAddress 用户博文存储地址
    BlogStorageAddress string = "./Blog/"
)

// MySQL参数配置
const (
    // DB 数据库类型
    DB string = "mysql"
    // DBUserName 数据库用户名称
    DBUserName string = "mlp"
    // DBPassword 数据库密码
    DBPassword string = "1984051718"
    // DBName 数据库名称
    DBName string = "mlp"
    // DBRemote 数据库占用端口
    DBRemote string = "127.0.0.1:3306"
    // DBCharset 数据库编码
    DBCharset string = "utf8"
    // DBParseTime 解析时间
    DBParseTime string = "True"
    // DBLoc 本地用户
    DBLoc string = "Local"
)

// Redis参数配置
const (
    // RedisRemote 数据库占用端口
    RedisRemote string = "localhost:6379"
    // RedisPassword 数据库密码
    RedisPassword string = "1984051718"
    // RedisName 数据库名称
    RedisName int = 0
)

// Mongo参数配置
const (
    // MongoAddress 数据库地址
    MongoAddress string = "mongodb://localhost:27017"
    // MongoDBName 数据库名称
    MongoDBName string = "mlp"
    // MongoTabName 数据表名称
    MongoTabName string = "blog"
)
