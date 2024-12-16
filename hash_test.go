package crypto

import (
	"fmt"
	"os"
	"testing"
)

func TestMd5String(t *testing.T) {
	md5 := Md5String([]byte("abc"))
	if len(md5) != 32 {
		t.Error(md5)
	}

	t.Log(md5)
}

func TestMd5File(t *testing.T) {
	path := fmt.Sprintf("%s/devspace/goinbox/crypto/hash.go", os.Getenv("HOME"))
	md5, err := Md5FileByPath(path)
	if err != nil {
		t.Error(err)
	}

	t.Log(md5, err)
}

func TestSha1String(t *testing.T) {
	sha1 := Sha1String([]byte("123"))
	if len(sha1) != 40 {
		t.Error(sha1)
	}
	t.Log(sha1)
}

func TestSha256String(t *testing.T) {
	sha256 := Sha256String([]byte("abczyx"))
	if len(sha256) != 64 {
		t.Error(sha256)
	}
	t.Log(sha256)
}
