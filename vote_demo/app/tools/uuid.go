package tools

import (
	"fmt"

	"github.com/google/uuid"
)

func GetUUID() string {
	id := uuid.New() // 默认v4版本 基于随机数
	fmt.Printf("uuid:%s, version:%d\n", id.String(), id.Version())
	return id.String()

}
