package utils

import "github.com/gofrs/uuid"

// GetUUID 根据 idtype 生成32位随机数
func GetUUID(idtype string) string {
    uid := uuid.NewV5(uuid.Must(uuid.NewV4()), idtype).String()
    uid = uid[:8] + uid[9:13] + uid[14:18] + uid[19:23] + uid[24:]
    return uid
}
