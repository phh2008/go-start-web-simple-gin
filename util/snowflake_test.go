package util

import (
	"fmt"
	"testing"
)

func TestIdWork(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println(IdWork.Generate().Int64())
	}
}
