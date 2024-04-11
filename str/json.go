package str

import "encoding/json"

func ToJson(v interface{}) (string, error) {
	if v == nil {
		return "", nil
	}
	d, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(d), nil
}
