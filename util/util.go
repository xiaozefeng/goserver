package util

import "github.com/teris-io/shortid"

func GetSorterId() string {
	r, _ := shortid.Generate()
	return r
}


