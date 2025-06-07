package utils

import (
	"context"
	"fmt"
)

type myLovelyKey string

func NewContextWithValue(parent context.Context, value any, key string) context.Context {
	ctx := context.WithValue(parent, myLovelyKey(key), value)
	return ctx
}

func GetValueFromContext[T any](ctx context.Context, key string) (*T, error) {
	if val := ctx.Value(myLovelyKey(key)); val != nil {
		res, ok := val.(T)
		if !ok {
			return nil, fmt.Errorf("casting the value to the given type failed")
		}
		return &res, nil
	}

	return nil, fmt.Errorf("the data with the key %v not found", key)
}
