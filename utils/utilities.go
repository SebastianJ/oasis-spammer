package utils

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"math/rand"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

const (
	Wei  = 1e18
	Gwei = 1e9
)

// FetchReceivers - fetch a list of proxies from a specified file
func FetchReceivers(filePath string) (lines []string, err error) {
	data, err := ReadFileToString(filePath)

	if err != nil {
		return nil, err
	}

	if len(data) > 0 {
		lines = strings.Split(string(data), "\n")

		if strings.Contains(data, "\n") {
			lines = lines[:len(lines)-1]
		}
	}

	return lines, nil
}

// RandomStringSliceItem - fetches a random string from a given string slice
func RandomStringSliceItem(r *rand.Rand, items []string) string {
	return items[r.Intn(len(items))]
}

// ReadFileToString - check if a file exists, proceed to read it to memory if it does
func ReadFileToString(filePath string) (string, error) {
	if FileExists(filePath) {
		data, err := ioutil.ReadFile(filePath)

		if err != nil {
			return "", err
		}

		return string(data), nil
	} else {
		return "", fmt.Errorf("file %s doesn't exist - make sure it exists or that you've specified the correct path for it", filePath)
	}
}

// FileExists - checks if a given file exists
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// ParseYaml - parses yaml into a specific type
func ParseYaml(path string, entity interface{}) error {
	yamlData, err := ReadFileToString(path)

	if err != nil {
		return err
	}

	err = yaml.Unmarshal([]byte(yamlData), entity)

	if err != nil {
		return err
	}

	return nil
}

// ConvertNumeralStringToBigInt - converts a numeral string back to a big float with the correct base set
func ConvertNumeralStringToBigInt(balance string) (*big.Int, error) {
	value := new(big.Float)
	value, ok := value.SetString(balance)

	if !ok {
		return nil, fmt.Errorf("can't convert balance string %s to a float balance", balance)
	}

	base := new(big.Float).SetInt(big.NewInt(Wei))
	value.Mul(value, base)

	result := new(big.Int)
	uintval, _ := value.Uint64()
	result.SetUint64(uintval)

	return result, nil
}
