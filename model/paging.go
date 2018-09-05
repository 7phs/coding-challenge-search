package model

import (
	"crypto/md5"
	"fmt"
)

type Paging struct {
	Start int `json:"start"`
	Limit int `json:"limit"`
}

func (o Paging) Hash(b []byte) []byte {
	hash := md5.New()

	if len(b) > 0 {
		hash.Write(b)
	}

	return hash.Sum([]byte(fmt.Sprint(o.Start, ";", o.Limit)))
}
