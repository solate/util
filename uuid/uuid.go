package uuid

import "github.com/google/uuid"

// 获取uuid
func GetUuid() (id string) {
	return uuid.New().String()
}
