package param

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func LookupEnvInt64(key string) (int64, error) {
	if val, ok := os.LookupEnv(key); ok {
		return strconv.ParseInt(val, 10, 64)
	}
	return 0, errors.New("not found")
}

func GetPriceFromConfig(path string) (int64, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("error while getting price", err)
		return 0, err
	}
	fmt.Println("data", string(data))
	return strconv.ParseInt(string(data), 10, 64)
}
