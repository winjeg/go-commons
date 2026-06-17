// copyleft anyone can use this code freely, and this repo will remain active
// welcome to come up with any kind of issues relating to this project

package conf

import (
	"os"
	"path/filepath"
)

// ReadConfigFile 读取配置文件内容
// 查找顺序：
// 1. 尝试直接打开文件
// 2. 在源文件目录及其父目录中搜索
// fileName: 配置文件名
// srcFile: 源文件路径（通过 runtime.Caller 获取）
// 返回文件内容，若文件不存在返回 nil
func ReadConfigFile(fileName, srcFile string) []byte {
	// 策略1：尝试直接打开文件
	if data, err := readFile(fileName); err == nil {
		return data
	}

	// 策略2：在源文件目录及其父目录中搜索
	dir := FindDirOfFile(fileName, srcFile)
	if dir == "" {
		return nil
	}

	configPath := filepath.Join(dir, fileName)
	data, err := readFile(configPath)
	if err != nil {
		return nil
	}

	return data
}
