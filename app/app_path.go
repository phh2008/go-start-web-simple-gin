package app

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

func GetAbsPath() string {
	dir, _ := GetAbsPathByExecutable()
	if strings.Contains(dir, GetTmpDir()) {
		return GetAbsPathByCaller()
	}
	return dir
}

// GetTmpDir 获取系统临时目录，兼容go run
func GetTmpDir() string {
	dir := os.Getenv("TEMP")
	if dir == "" {
		dir = os.Getenv("TMP")
	}
	res, _ := filepath.EvalSymlinks(dir)
	return res
}

// GetAbsPathByExecutable 获取当前执行文件绝对路径
func GetAbsPathByExecutable() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	res, err := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res, err
}

// GetAbsPathByCaller 获取当前执行文件绝对路径（go run）
func GetAbsPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}
