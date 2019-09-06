package util

import "github.com/teris-io/shortid"

func GetSorterId() (string, error) {
	return shortid.Generate()
}


