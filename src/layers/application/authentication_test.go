package application

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func Test_getSalt(t *testing.T) {
	v, _ := generate()
	fmt.Println(hex.EncodeToString(v))
	fmt.Println(v)
}
