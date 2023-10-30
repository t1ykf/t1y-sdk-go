package t1y

import (
	"crypto/md5"
	"fmt"
)

func MD5(input string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(input)))
}
