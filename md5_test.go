package crypto

import (
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
	md5, err := Md5File("/home/ligang/devspace/goinbox/crypto/md5.go")
	if err != nil {
		t.Error(err)
	}

	t.Log(md5, err)
}
