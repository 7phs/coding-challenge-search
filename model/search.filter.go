package model

import (
	"bytes"
	"crypto/md5"
	"strings"
)

const (
	KeywordsModeDefault KeywordsMode = 0
	KeywordsModeAnd     KeywordsMode = 1
	KeywordsModeOr      KeywordsMode = 2
)

type KeywordsMode int

func (o KeywordsMode) String() string {
	switch o {
	case KeywordsModeAnd:
		return "AND"
	case KeywordsModeOr:
		return "OR"
	default:
		return "AND"
	}
}

type SearchFilter struct {
	Mode     KeywordsMode   `json:"searchTerms,omitempty"`
	Keywords *SearchKeyword `json:"searchTerms"`
	Location Location       `json:"location"`
}

func (o *SearchFilter) Empty() bool {
	return o == nil || o.Keywords.Empty() && o.Location.Empty()
}

func (o *SearchFilter) String() string {
	if o == nil {
		return "nil"
	}

	return o.Mode.String() + "('" + o.Keywords.String() + "') + (" + o.Location.String() + ")"
}

func (o *SearchFilter) Hash(b []byte) []byte {
	if o == nil {
		return nil
	}

	buf := bytes.NewBufferString(o.Mode.String())
	buf.WriteString(";")
	buf.WriteString(strings.Join(o.Keywords.Lemmas(), ">|<"))
	buf.WriteString(";")
	buf.WriteString(o.Location.String())

	hash := md5.New()

	if len(b) > 0 {
		hash.Write(b)
	}

	return hash.Sum(buf.Bytes())
}
