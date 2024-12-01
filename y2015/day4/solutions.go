package day4

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

func Part1(input *string) string {
	for i := 0; ; i++ {
		hash := md5.Sum([]byte(*input + strconv.Itoa(i)))
		if strings.HasPrefix(fmt.Sprintf("%x", hash), "00000") {
			return strconv.Itoa(i)
		}
	}
}

func Part2(input *string) string {
	for i := 0; ; i++ {
		hash := md5.Sum([]byte(*input + strconv.Itoa(i)))
		if strings.HasPrefix(fmt.Sprintf("%x", hash), "000000") {
			return strconv.Itoa(i)
		}
	}
}
