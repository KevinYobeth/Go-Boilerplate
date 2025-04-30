package utils

import (
	"runtime"
	"strings"
)

func GetFnCallerName(skip, indexGet int) string {
	skip++

	pc, _, _, ok := runtime.Caller(skip)
	if !ok {
		return ""
	}

	fullName := runtime.FuncForPC(pc).Name()
	funcName := strings.Split(fullName, ".")

	if indexGet >= len(funcName) {
		return fullName
	}

	partitionGet := len(funcName) - (len(funcName) - indexGet)

	return strings.Join(funcName[partitionGet:], ".")
}
