package param

import (
	"errors"
	"os"
	"strconv"
)

func LookupEnvInt64(key string) (int64, error) {
	if val, ok := os.LookupEnv(key); ok {
		return strconv.ParseInt(val, 10, 64)
	}
	return 0, errors.New("not found")
}
