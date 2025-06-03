package models

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func GetGravatarURL(email string, size int) string {
	email = strings.ToLower(strings.TrimSpace(email))
	hash := md5.Sum([]byte(email))
	return fmt.Sprintf("https://www.gravatar.com/avatar/%x?s=%d&d=identicon", hash, size)
}
