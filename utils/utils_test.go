package utils

import (
	"strconv"
	"testing"
)

func TestGetRootPackageOffset(t *testing.T) {
	base := "/Users/tlatins/gocode/src/github.com/tlatin/"
	offset := GetRootPackageOffset(base + rootPackageName)
	if 0 != offset {
		t.Error("Root Package returned the wrong offset. Expected: 2, Saw: " + strconv.Itoa(offset))
	}
	offset = GetRootPackageOffset(base + rootPackageName + "/controller/cron")
	if 2 != offset {
		t.Error("Root Package returned the wrong offset. Expected: 2, Saw: " + strconv.Itoa(offset))
	}
}
