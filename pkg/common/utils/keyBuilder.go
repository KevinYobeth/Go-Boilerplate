package utils

import (
	"fmt"
	"strings"
)

func NewRedisKey(key ...string) string {
	return fmt.Sprintf("boilerplate:%s", strings.Join(key, ":"))
}
