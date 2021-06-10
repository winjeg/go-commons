package str

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"github.com/satori/go.uuid"
)

const (
	// numbers only
	KindNumber = 0
	// lower alphabets
	KindLower = 1
	// upper alphabets
	KindUpper = 2
	// numbers and alphabets
	KindAll = 3
	// all kinds with special chars
	KindAllWithSpecial = 4
)

// export RandStr
// random strings
func RandStr(size int, kind int) []byte {
	kinds := [][]int{{10, 48}, {26, 97}, {26, 65}}
	specialChars := []byte{95, 45, 46, 35, 36, 37, 38}
	specialCharLen := len(specialChars)
	iKind, result := kind, make([]byte, size)
	isAll := kind == 3
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {

		// random iKind
		if isAll {
			iKind = rand.Intn(3)
		}
		if kind == KindAllWithSpecial {
			iKind = rand.Intn(4)
		}
		if iKind == 3 {
			result[i] = specialChars[rand.Intn(specialCharLen)]
		} else {
			scope, base := kinds[iKind][0], kinds[iKind][1]
			result[i] = uint8(base + rand.Intn(scope))
		}
	}
	return result
}

// RandomNumAlphabets
// random number and alphabets
func RandomNumAlphabets(length int) string {
	return string(RandStr(length, KindAll))
}

// export RandomNumbers
// random numbers
func RandomNumbers(length int) string {
	return string(RandStr(length, KindNumber))
}

// export RandomAlphabetsLower
// random alphabets in lower case
func RandomAlphabetsLower(length int) string {
	return string(RandStr(length, KindLower))
}

// export RandomAlphabetsUpper
// random alphabets in upper case
func RandomAlphabetsUpper(length int) string {
	return string(RandStr(length, KindUpper))
}

// export RandomStrWithSpecialChars
// random string with special chars including _-.#$%&
func RandomStrWithSpecialChars(length int) string {
	return string(RandStr(length, KindAllWithSpecial))
}

// export UUID
// is to generate unique ids
func UUID() string {
	uid := uuid.NewV4()
	return uid.String()
}

// export UUIDShort
// this method will generate a unique id using uuid, but the result is too long
// so we just use the digits from 0 to 8, thus, increasing the possibility to get a
// duplicated id, but It's okay
// not true uuid, not for tons of ids
func UUIDShort() string {
	u2 := uuid.NewV4()
	d := u2.String()
	return d[24:] + d[9:13]
}

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

// export IsAnyBlank
// judge if any strings is blank
func IsAnyBlank(strArr ...string) bool {
	if strArr == nil {
		return true
	}
	for _, str := range strArr {
		if IsBlank(str) {
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

// remove the leading comma and trailing comma
func TrimComma(str string) string {
	return Trim(str, ",")
}

// remove leading and trailing dot
func TrimDot(str string) string {
	return Trim(str, ".")
}

// trim something down, not a cut set. which is much different from strings.Trim
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

// same as the strings.Join
func Join(strs []string, j string) string {
	return strings.Join(strs, j)
}

// convert a string to a byte
func ToBytes(str string) []byte {
	bs := (*[2]uintptr)(unsafe.Pointer(&str))
	b := [3]uintptr{bs[0], bs[1], bs[1]}
	return *(*[]byte)(unsafe.Pointer(&b))
}

// convert a byte array to a string
func FromBytes(data []byte) string {
	return *(*string)(unsafe.Pointer(&data))
}

// same as the strings.Join
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

// split strings
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

// join args if the args is not empty
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
