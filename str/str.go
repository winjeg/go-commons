package str

import (
	"github.com/satori/go.uuid"
)

// generate a random string using number, alphabets and some symbols
func RandomString(length int) string {
	// TODO
	return ""
}

func RandomNumers(length int) string {
	// TODO
	return ""
}

func RandomAlphabets(length int) string {
	// TODO
	return ""
}


// generate uuid
func UUID() string {
	uid := uuid.NewV4()
	return uid.String()
}

// this method will generate a unique id using uuid, but the result is too long
// so we just use the digits from 0 to 8, thus, increasing the possibility to get a
// duplicated id, but It's okay
// not true uuid, not for tons of ids
func UUIDShort() string {
	u2 := uuid.NewV4()
	d := u2.String()
	return d[24:] + d[9:13]
}
