package crypto

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"

	"fmt"
	"io"
	"os"
)

func Md5String(data []byte) string {
	return fmt.Sprintf("%x", md5.Sum(data))
}

func Md5File(file io.Reader) (string, error) {
	h := md5.New()
	if _, err := io.Copy(h, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func Md5FileByPath(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = f.Close()
	}()

	return Md5File(f)
}

func Sha1String(data []byte) string {
	return fmt.Sprintf("%x", sha1.Sum(data))
}

func Sha256String(data []byte) string {
	return fmt.Sprintf("%x", sha256.Sum256(data))
}
