package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	number := int64(1234)
	bs := strconv.FormatInt(number, 2)
	strings.Count(strconv.FormatInt(number, 2), "1")
	fmt.Println(bs)
	fmt.Println(strings.Count(bs, "1"))
}
