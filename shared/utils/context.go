package utils

import (
	"context"
	"fmt"
	"reflect"
	"unsafe"

	"github.com/kevinyobeth/go-boilerplate/shared/constants"
)

func PrintContextValues(ctx interface{}, inner bool) {
	contextValues := reflect.ValueOf(ctx).Elem()
	contextKeys := reflect.TypeOf(ctx).Elem()

	if !inner {
		fmt.Printf("\nFields for %s.%s\n", contextKeys.PkgPath(), contextKeys.Name())
	}

	if contextKeys.Kind() == reflect.Struct {
		for i := 0; i < contextValues.NumField(); i++ {
			reflectValue := contextValues.Field(i)
			reflectValue = reflect.NewAt(reflectValue.Type(), unsafe.Pointer(reflectValue.UnsafeAddr())).Elem()

			reflectField := contextKeys.Field(i)

			if reflectField.Name == "Context" {
				PrintContextValues(reflectValue.Interface(), true)
			} else {
				fmt.Printf("field name: %+v\n", reflectField.Name)
				fmt.Printf("value: %+v\n", reflectValue.Interface())
			}
		}
	} else {
		fmt.Printf("context is empty (int)\n")
	}
}

func AddToCtx(ctx context.Context, key constants.ContextKey, value interface{}) context.Context {
	return context.WithValue(ctx, key, value)
}

func ReadFromCtx(ctx context.Context, key constants.ContextKey) interface{} {
	return ctx.Value(key)
}
