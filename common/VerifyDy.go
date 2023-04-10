package common

import (
	"crypto/md5"
	"encoding/base64"
	"sort"
	"strings"
	"xyx/config"
)

// 验证抖音前面
func Signature(header map[string]string) string {
	bodyStr := config.BodyStr
	secret := config.Secret

	keyList := make([]string, 0, 4)
	for key, _ := range header {
		keyList = append(keyList, key)
	}
	sort.Slice(keyList, func(i, j int) bool {
		return keyList[i] < keyList[j]
	})
	kvList := make([]string, 0, 4)
	for _, key := range keyList {
		kvList = append(kvList, key+"="+header[key])
	}
	urlParams := strings.Join(kvList, "&")
	rawData := urlParams + bodyStr + secret
	md5Result := md5.Sum([]byte(rawData))
	return base64.StdEncoding.EncodeToString(md5Result[:])
}
