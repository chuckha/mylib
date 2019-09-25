package mylib

import (
	"github.com/chuckha/my-repo"
)

func IgnoreError() string {
	s, err := my_repo.MyRepo()
	if err != nil {
		return ""
	}
	return s
}
