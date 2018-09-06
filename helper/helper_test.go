package helper

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"testing"

	"github.com/bmizerany/assert"
)

func TestGetEnv(t *testing.T) {
	expected := map[string]interface{}{
		"TEST_HELPER_1": "value1",
		"TEST_HELPER_2": int64(15),
		"TEST_HELPER_3": "a15",
		"TEST_HELPER_4": true,
	}

	defer func() func() {
		type prevEnv struct {
			value string
			ok    bool
		}

		prevValue := make(map[string]prevEnv)

		for key, value := range expected {
			v, ok := os.LookupEnv(key)

			prevValue[key] = prevEnv{
				value: v,
				ok:    ok,
			}

			os.Setenv(key, fmt.Sprint(value))
		}

		return func() {
			for key, prev := range prevValue {
				if prev.ok {
					os.Setenv(key, prev.value)
				} else {
					os.Setenv(key, "")
				}
			}
		}
	}()()

	unknownName := "UNKNOWN" + strconv.Itoa(rand.Int())

	name := "TEST_HELPER_1"
	assert.Equal(t, expected[name], GetEnvStr(name, ""), name)

	expectedStr := "unknown-default"
	assert.Equal(t, expectedStr, GetEnvStr(unknownName, expectedStr), name)

	name = "TEST_HELPER_2"
	assert.Equal(t, expected[name], GetEnvInt64(name, 0), name)

	expectedInt64 := 2 + rand.Int63n(300000)
	assert.Equal(t, expectedInt64, GetEnvInt64(unknownName, expectedInt64), name)

	name = "TEST_HELPER_3"
	expectedInt64 = 2 + rand.Int63n(300000)
	assert.Equal(t, expectedInt64, GetEnvInt64(name, expectedInt64), name)

	name = "TEST_HELPER_4"
	assert.Equal(t, true, GetEnvBool(name, false), name)

	assert.Equal(t, true, GetEnvBool(unknownName, true), name)
}
