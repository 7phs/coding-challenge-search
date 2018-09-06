package common

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorRecordList_Error(t *testing.T) {
	var (
		exist     ErrorRecordList
		existResp RespError
		expected  = ErrorRecordList{
			{
				Id:   1,
				Desc: "err-true",
			},
			{
				Id:   2,
				Desc: "err-15",
			},
			{
				Id:   3,
				Desc: "3err-1",
			},
			{
				Id:   3,
				Desc: "3err-2",
			},
			{
				Id:   4,
				Desc: "5err-1",
			},
			{
				Id:   5,
				Desc: "5err-1",
			},
		}
	)

	assert.Nil(t, exist, "init")
	assert.Empty(t, exist.Error(), "init")

	assert.Nil(t, existResp.Errors, "init")

	exist.AddError(1, "err-", true)
	existResp.AddError(1, "err-", true)

	exist.AddErrorf(2, "err-%d", 15)
	existResp.AddErrorf(2, "err-%d", 15)

	errList := []error{
		errors.New("3err-1"),
		errors.New("3err-2"),
	}
	errRecList := ErrorRecordList{
		{
			Id:   4,
			Desc: "5err-1",
		},
		{
			Id:   5,
			Desc: "5err-1",
		},
	}

	exist.AppendError(3, errList)
	existResp.AddError(3, errList[0].Error())
	existResp.AddError(3, errList[1].Error())

	exist.Append(errRecList)
	existResp.AppendError(errRecList)

	assert.Equal(t, expected, exist)
	assert.Equal(t, expected, existResp.Errors)

	expectedStr := "[1] err-true; [2] err-15; [3] 3err-1; [3] 3err-2; [4] 5err-1; [5] 5err-1"
	assert.Equal(t, expectedStr, exist.Error())
}

func TestListOfErr_Error(t *testing.T) {
	var (
		exist = ListOfErr{
			Id: 11,
		}
	)

	assert.Equal(t, false, exist.HasError(), "init")

	exist.Check(false, "err1")
	exist.Check(true, "err2")
	exist.Checkf(false, "err%d", 3)
	exist.Checkf(true, "err%d", 4)

	assert.Equal(t, true, exist.HasError())

	expectedStr := "[11] err2; [11] err4"
	assert.Equal(t, expectedStr, exist.Error())
}
