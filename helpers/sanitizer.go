package helpers

import (
	"strconv"
	"fmt"
)

func VarToFloat64(unk interface{}) (sanitizedInput float64) {
	sanitizedInput = float64(0)
	switch i := unk.(type) {
	case string:
		sanitizedInput, _ = strconv.ParseFloat(i, 64)
	case float64:
		sanitizedInput = i
	case float32:
		sanitizedInput = float64(i)
	case int64:
		sanitizedInput = float64(i)
	case int32:
		sanitizedInput = float64(i)
	case int:
		sanitizedInput = float64(i)
	case uint64:
		sanitizedInput = float64(i)
	case uint32:
		sanitizedInput = float64(i)
	case uint:
		sanitizedInput = float64(i)
	}
	return
}

func VarToFloat32(unk interface{}) (sanitizedInput float32) {
	sanitizedInput = float32(0)
	switch i := unk.(type) {
	case string:
		tempInput, _ := strconv.ParseFloat(i, 64)
		sanitizedInput = float32(tempInput)
	case float64:
		sanitizedInput = float32(i)
	case float32:
		sanitizedInput = i
	case int64:
		sanitizedInput = float32(i)
	case int32:
		sanitizedInput = float32(i)
	case int:
		sanitizedInput = float32(i)
	case uint64:
		sanitizedInput = float32(i)
	case uint32:
		sanitizedInput = float32(i)
	case uint:
		sanitizedInput = float32(i)
	}
	return
}

func VarToInt64(unk interface{}) (sanitizedInput int64) {
	sanitizedInput = int64(0)
	switch i := unk.(type) {
	case string:
		tempInput, _ := strconv.ParseFloat(i, 64)
		sanitizedInput = int64(tempInput)
	case float64:
		sanitizedInput = int64(i)
	case float32:
		sanitizedInput = int64(i)
	case int64:
		sanitizedInput = i
	case int32:
		sanitizedInput = int64(i)
	case int:
		sanitizedInput = int64(i)
	case uint64:
		sanitizedInput = int64(i)
	case uint32:
		sanitizedInput = int64(i)
	case uint:
		sanitizedInput = int64(i)
	}
	return
}

func VarToInt32(unk interface{}) (sanitizedInput int32) {
	sanitizedInput = int32(0)
	switch i := unk.(type) {
	case string:
		tempInput, _ := strconv.ParseFloat(i, 64)
		sanitizedInput = int32(tempInput)
	case float64:
		sanitizedInput = int32(i)
	case float32:
		sanitizedInput = int32(i)
	case int64:
		sanitizedInput = int32(i)
	case int32:
		sanitizedInput = i
	case int:
		sanitizedInput = int32(i)
	case uint64:
		sanitizedInput = int32(i)
	case uint32:
		sanitizedInput = int32(i)
	case uint:
		sanitizedInput = int32(i)
	}
	return
}

func VarToString(unk interface{}) (sanitizedInput string) {
	sanitizedInput = ""
	switch i := unk.(type) {
	case string:
		sanitizedInput = i
	case float64:
		sanitizedInput = fmt.Sprintf("%g", i)
	case float32:
		sanitizedInput = fmt.Sprintf("%g", i)
	case int64:
		sanitizedInput = strconv.FormatInt(i, 10)
	case int32:
		sanitizedInput = strconv.FormatInt(int64(i), 10)
	case int:
		sanitizedInput = strconv.FormatInt(int64(i), 10)
	case uint64:
		sanitizedInput = strconv.FormatInt(int64(i), 10)
	case uint32:
		sanitizedInput = strconv.FormatInt(int64(i), 10)
	case uint:
		sanitizedInput = strconv.FormatInt(int64(i), 10)
	}
	return
}

func DeleteEmptyMapValues(s map[string]string) map[string]string {
	r := make(map[string]string)
	for index, str := range s {
		if str != "" {
			r[index] = str
		}
	}
	return r
}
