package config

// 基本参数配置
const (
    // BackendAddress 后端地址
    BackendAddress string = "localhost"
    // Port 端口
    Port string = ":8001"
    // FileSizeLimit 上传文件最大限制
    FileSizeLimit int64 = 1024
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
