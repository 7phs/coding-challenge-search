package helper

import (
	"bytes"
	"errors"
	"net"
	"net/url"
	"os"
	"strings"
)

// getEnv return env variable or default value provided
func GetEnvStr(name, defaultV string) string {
	if value, ok := os.LookupEnv(name); ok {
		return value
	}

	return defaultV
}

func GetEnvBool(name string, defaultV bool) bool {
	if value, ok := os.LookupEnv(name); ok {
		return strings.ToLower(value) == "true"
	}

	return defaultV
}

func ErrorFromList(errList []error) error {
	buf := bytes.NewBufferString("")

	for i, err := range errList {
		if i > 0 {
			buf.WriteString("; ")
		}

		buf.WriteString(err.Error())
	}

	return errors.New(buf.String())
}

func ValidateHostUrl(u string) error {
	u = strings.TrimSpace(u)
	if len(u) == 0 {
		return errors.New("empty")
	} else if u, err := url.Parse(u); err != nil {
		return errors.New("invalid: " + err.Error())
	} else {
		if strings.Index(u.Host, ":") > 0 {
			u.Host, _, _ = net.SplitHostPort(u.Host)
		}
		if _, err := net.LookupHost(u.Host); err != nil {
			return errors.New("resolve err: " + err.Error())
		}
	}

	return nil
}
