package str

import (
	"bytes"
	"strconv"
	"strings"
	"unsafe"
)

// export IsEmpty
// judge if a string is empty
func IsEmpty(str string) bool {
	return len(str) == 0
}

// export IsBlank
// judge if a string is blank
func IsBlank(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}

// export IsNotBlank
// is not blank.
func IsNotBlank(str string) bool {
	return !IsBlank(str)
}

// export IsAllBlank
// judge if all strings is blank
func IsAllBlank(strArr ...string) bool {
	for _, str := range strArr {
		if !IsBlank(str) {
			return false
		}
	}
	return true
}

// export IsNoneBlank
// to judge if strings is none blank
func IsNoneBlank(strArr ...string) bool {
	if strArr == nil {
		return false
	}
	for _, v := range strArr {
		if IsBlank(v) {
			return false
		}
	}
	return true
}

// export HasAnyBlank
// if the given parameters has any blank, true value will be returned
func HasAnyBlank(strArr ...string) bool {
	if strArr == nil {
		return true
	}
	for _, v := range strArr {
		if IsBlank(v) {
			return true
		}
	}
	return false
}

// export NotAllBlank
// not all strings in a given array is blank
func NotAllBlank(strArr ...string) bool {
	if strArr == nil {
		return false
	}
	for _, v := range strArr {
		if IsNotBlank(v) {
			return true
		}
	}
	return false
}

// export ReplaceAll
// replace all old str to new str for a string
func ReplaceAll(src, old, new string) string {
	return strings.Replace(src, old, new, -1)
}

// TrimComma remove the leading comma and trailing comma
func TrimComma(str string) string {
	return Trim(str, ",")
}

// TrimDot remove leading and trailing dot
func TrimDot(str string) string {
	return Trim(str, ".")
}

// Trim trim something down, not a cut set. which is much different from strings.Trim
// trim space first then the trim param set
// it only trims the first occurrences of the string to be trimmed
func Trim(str, trim string) string {
	str = strings.TrimSpace(str)
	if strings.EqualFold(str[len(str)-len(trim):], trim) {
		str = str[:len(str)-len(trim)]
	}
	if strings.EqualFold(str[0:len(trim)], trim) {
		str = str[len(trim):]
	}
	return str
}

// Join same as the strings.Join
func Join(strArr []string, j string) string {
	return strings.Join(strArr, j)
}

// JoinInt same as the strings.Join
func JoinInt(arr []int, j string) string {
	if len(arr) == 0 {
		return ""
	}
	var result string
	for _, v := range arr {
		result += strconv.Itoa(v) + j
	}
	return result[:len(result)-1]
}

// SplitInt split strings
func SplitInt(ori, sep string) ([]int, error) {
	if len(ori) == 0 {
		return nil, nil
	}
	arr := strings.Split(ori, sep)
	result := make([]int, 0, len(arr))
	for _, v := range arr {
		i, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		result = append(result, i)
	}
	return result, nil
}

func SplitInt64(ori, sep string) ([]int64, error) {
	if len(ori) == 0 {
		return nil, nil
	}
	arr := strings.Split(ori, sep)
	result := make([]int64, 0, len(arr))
	for _, v := range arr {
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, err
		}
		result = append(result, i)
	}
	return result, nil
}

// JoinIfNotEmpty join args if the args is not empty
func JoinIfNotEmpty(sep string, args ...string) string {
	if len(args) < 1 {
		return ""
	}
	newArr := make([]string, 0, len(args))
	for i := range args {
		if len(args[i]) > 0 {
			newArr = append(newArr, args[i])
		}
	}
	var result string
	for i := range newArr {
		if i == len(newArr)-1 {
			result += newArr[i]
		} else {
			result += newArr[i] + sep
		}
	}
	return result
}

// export Repeat
// repeat the string with given times
func Repeat(str string, n int) string {
	if len(str) == 0 || n < 2 {
		return str
	}
	var buffer bytes.Buffer
	for i := 0; i < n; i++ {
		buffer.WriteString(str)
	}
	return buffer.String()
}

// export JoinStr
// join given strings together
func JoinStr(strArr ...string) string {
	if len(strArr) == 0 {
		return ""
	}
	var buffer bytes.Buffer
	for _, v := range strArr {
		buffer.WriteString(v)
	}
	return buffer.String()
}

// ToBytes convert a string to a byte
func ToBytes(str string) []byte {
	bs := (*[2]uintptr)(unsafe.Pointer(&str))
	b := [3]uintptr{bs[0], bs[1], bs[1]}
	return *(*[]byte)(unsafe.Pointer(&b))
}

// FromBytes convert a byte array to a string
func FromBytes(data []byte) string {
	return *(*string)(unsafe.Pointer(&data))
}

// StartsWith to judge if the given str starts with
// the part given or not
func StartsWith(str, part string) bool {
	if len(part) == 0 {
		return true
	}
	index := strings.Index(str, part)
	return index == 0
}

// EndsWith to judge if the given str ends with
// the part given or not
func EndsWith(str, part string) bool {
	if len(part) == 0 {
		return true
	}
	lastIndex := strings.LastIndex(str, part)
	return len(str) == len(part)+lastIndex
}
