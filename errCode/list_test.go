package errCode

import (
	"errors"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrList_Error(t *testing.T) {
	var (
		exist  ErrList
		expStr []string
	)

	assert.Nil(t, exist.Result(), "init")
	assert.Empty(t, exist.Error(), "init")

	for i := 0; i < 5; i++ {
		exist.Add(
			errors.New("e"+strconv.Itoa(i)),
			errors.New("v"+strconv.Itoa(i)))

		expStr = append(expStr, "e"+strconv.Itoa(i))
		expStr = append(expStr, "v"+strconv.Itoa(i))
	}

	assert.NotNil(t, exist.Result(), "after init")

	assert.Equal(t, strings.Join(expStr, "; "), exist.Error())
}
