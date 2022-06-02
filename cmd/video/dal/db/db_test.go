package db

import (
	"context"
	"fmt"
	"testing"
)

func MGetFavoriteListTest(t *testing.T) {
	res, err := MGetFavoriteList(context.Background(), 7)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(res)
}
