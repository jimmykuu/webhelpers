package webhelpers

import (
	"crypto/md5"
	"fmt"
	"io"
)

func Gravatar(email string, size uint16) string {
	h := md5.New()
	io.WriteString(h, email)
	return fmt.Sprintf("http://www.gravatar.com/avatar/%x?s=%d", h.Sum(nil), size)
}
