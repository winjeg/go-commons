package conf

import (
	"os"
	"path/filepath"
)

// getFileDirFromSrc 从源代码所在目录向上搜索配置文件
// 不改变工作目录，使用绝对路径操作
// fileName: 要查找的文件名
// srcFile: 源文件的路径
// 返回找到配置文件的目录，若未找到返回空字符串
func getFileDirFromSrc(fileName, srcFile string) string {
	// 获取源文件的绝对路径
	dir, err := filepath.Abs(srcFile)
	if err != nil {
		return ""
	}

	// 如果srcFile指向文件，获取其所在目录
	if info, err := os.Stat(dir); err == nil && !info.IsDir() {
		dir = filepath.Dir(dir)
	}

	// 从源文件所在目录向上逐层搜索
	currentDir := dir
	for {
		// 检查当前目录是否包含配置文件
		configPath := filepath.Join(currentDir, fileName)
		if _, err := os.Stat(configPath); err == nil {
			return currentDir
		}

		// 获取父目录
		parentDir := filepath.Dir(currentDir)

		// 如果已经到达根目录，停止搜索
		if parentDir == currentDir {
			// 已到达文件系统根目录
			return ""
		}

		currentDir = parentDir
	}
}

// getFileDirFromExecutable 从可执行文件所在目录查找配置文件
// fileName: 要查找的文件名
// 返回找到配置文件的目录，若未找到返回空字符串
func getFileDirFromExecutable(fileName string) string {
	execPath, err := os.Executable()
	if err != nil {
		return ""
	}

	dir := filepath.Dir(execPath)
	configPath := filepath.Join(dir, fileName)

	if _, err := os.Stat(configPath); err == nil {
		return dir
	}

	return ""
}

// FindDirOfFile 查找配置文件所在目录
// 优先从可执行文件目录查找，然后从源代码目录向上查找
// fileName: 配置文件名 (e.g., "config.yaml")
// srcFile: 调用者的源文件路径 (通过 runtime.Caller 获得)
// 返回找到配置文件的目录，若未找到返回空字符串
func FindDirOfFile(fileName, srcFile string) string {
	// 策略1：从可执行文件所在目录查找
	if dir := getFileDirFromExecutable(fileName); dir != "" {
		return dir
	}

	// 策略2：从源代码所在目录向上查找
	return getFileDirFromSrc(fileName, srcFile)
}
