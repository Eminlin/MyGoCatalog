package uuid

import (
	"strings"

	"github.com/google/uuid"
)

// GenUUID 生成一个随机的唯一ID
func GenUUID() string {
	return uuid.New().String()
}

// GenUUID16 截取uuid前16位
func GenUUID16() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")[0:16]
}

// ParseUUIDFromStr 从指定的字符串生成uuid并原样返回该uuid字符串
// 必须符合uuid格式，否则返回一个error
func ParseUUIDFromStr(str string) (string, error) {
	u, err := uuid.Parse(str)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}
