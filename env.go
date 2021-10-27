package env

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

func Getenv(key string, apply func(string) (string, error)) (ret string, err error) {
	ret = os.Getenv(key)
	if apply != nil {
		ret, err = apply(ret)
	}
	return

}
func GetEnvList(env string, sep string) []string {
	v := os.Getenv(env)
	sv := strings.Split(v, sep)

	ret := []string{}
	for _, s := range sv {
		ss := strings.TrimSpace(s)
		if len(ss) > 0 {
			ret = append(ret, ss)
		}
	}
	return ret
}

func MustGetenv(key string) string {
	ret, _ := Getenv(key, func(s string) (string, error) {
		if s == "" {
			panic(fmt.Errorf("env variable %s is empty", key))
		}
		return s, nil

	})
	return ret
}

func DecodeBase64(s string) (string, error) {
	res, err := base64.RawStdEncoding.DecodeString(s)
	return string(res), err
}
