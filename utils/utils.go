package utils

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func HexNumberToString(hexaString string) string {
	// replace 0x or 0X with empty String
	numberStr := strings.Replace(hexaString, "0x", "", -1)
	numberStr = strings.Replace(numberStr, "0X", "", -1)
	return numberStr
}

func GetHexToUint64(data []byte) (uint64, error) {
	var val string
	err := json.Unmarshal(data, &val)
	if err != nil {
		return 0, fmt.Errorf("json.Unmarshal:%s error:%s", data, err)
	}
	val = HexNumberToString(val)
	intNumber, err := strconv.ParseUint(val, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("stringVal:%s error:%s", val, err)
	}
	return intNumber, nil
}
