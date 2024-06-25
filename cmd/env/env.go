package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for _, e := range os.Environ() {
		s := strings.Split(e, "=")
		key, value := s[0], s[1]
		fmt.Println(key, "=", value)
	}
}
