package model

import (
	"crypto/md5"
	"fmt"
	"os"
)

type Paging struct {
	Start int `json:"start" validation:"min:0"`
	Limit int `json:"limit" validation:"min:0"`
}

func (o Paging) Hash(b []byte) []byte {
	hash := md5.New()

	if len(b) > 0 {
		hash.Write(b)
	}

	return hash.Sum([]byte(fmt.Sprint(o.Start, ";", o.Limit)))
}

func (o Paging) StartLimit(ln int) (int, int, error) {
	if o.Start > ln {
		return 0, 0, os.ErrInvalid
	}

	start := o.Start
	limit := o.Limit
	if start+limit > ln {
		limit = ln - start
	}

	return start, start + limit, nil
}
