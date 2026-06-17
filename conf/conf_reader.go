// copyleft anyone can use this code freely, and this repo will remain active
// welcome to come up with any kind of issues relating to this project

package conf

import (
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"gopkg.in/ini.v1"
	"gopkg.in/yaml.v2"
)

// Yaml2Object 将YAML配置文件解析为对象
// 支持三种查找方式：
// 1. 直接打开文件（若fileName是绝对路径或相对当前目录存在）
// 2. 从可执行文件目录向上查找
// 3. 从调用者源文件目录向上查找
func Yaml2Object(fileName string, object interface{}) error {
	data, err := readConfigFile(fileName)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, object)
}

// Ini2Object 将INI配置文件解析为对象
// 支持三种查找方式：
// 1. 直接打开文件（若fileName是绝对路径或相对当前目录存在）
// 2. 从可执行文件目录向上查找
// 3. 从调用者源文件目录向上查找
func Ini2Object(fileName string, object interface{}) error {
	data, err := readConfigFile(fileName)
	if err != nil {
		return err
	}
	return ini.MapTo(object, data)
}

// readConfigFile 内部函数：读取配置文件
// 查找顺序：
// 1. 尝试直接打开文件（支持绝对路径或相对路径）
// 2. 使用runtime.Caller获取调用者源文件，从该目录向上搜索
func readConfigFile(fileName string) ([]byte, error) {
	// 策略1：尝试直接打开文件（支持绝对路径或相对路径）
	if data, err := readFile(fileName); err == nil {
		return data, nil
	}

	// 策略2：使用runtime.Caller获取调用者信息，向上搜索配置文件
	// Caller(2) 跳过：readConfigFile -> Yaml2Object/Ini2Object -> 用户代码
	_, callerFile, _, ok := runtime.Caller(2)
	if !ok {
		return nil, errors.New("cannot get caller info")
	}

	dir := FindDirOfFile(fileName, callerFile)
	if dir == "" {
		return nil, errors.New("config file not found: " + fileName)
	}

	// 从找到的目录读取文件
	configPath := filepath.Join(dir, fileName)
	return readFile(configPath)
}

// readFile 读取文件内容
func readFile(filePath string) ([]byte, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeErr := f.Close(); closeErr != nil {
			log.Printf("error closing file: %v", closeErr)
		}
	}()

	return io.ReadAll(f)
}
