package props

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

var (
	regex = regexp.MustCompile(`^([^=]+)=(.*)$`)
)

type Properties map[string]string

func parseFile(in io.Reader) (Properties, error) {
	scanner := bufio.NewScanner(in)
	props := Properties{}
	lineNumber := 1
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		// Skip comment
		if len(line) == 0 || line[0] == '#' {
			lineNumber++
			continue
		}
		if groups := regex.FindStringSubmatch(line); groups != nil {
			key, value := groups[1], groups[2]
			props[strings.TrimSpace(key)] = strings.TrimSpace(value)
		} else {
			return nil, fmt.Errorf("invalid syntax at line %d", lineNumber)
		}
		lineNumber++
	}
	err := scanner.Err()
	if err != nil {
		return nil, err
	}
	return props, nil
}

func LoadFile(fileName string) (Properties, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	return parseFile(file)
}

func FromString(content string) (Properties) {
	if len(content) < 1 {
		return map[string]string{}
	}
	result := make(map[string]string, 100)
	lines := strings.Split(content, "\n")
	for _, v := range lines {
		l := strings.TrimSpace(v)
		if len(v) < 1 || v[0] == '#' {
			continue
		}
		idx := strings.Index(l, "=")
		if idx == -1 {
			continue
		}
		result[l[0:idx]] = l[idx+1:]
	}
	return result
}

func (p Properties) String() string {
	result := ""
	for k, v := range p {
		result += k + "=" + v + "\n"
	}
	return result
}
