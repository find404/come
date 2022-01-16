package librarys

/**
 * 当前包公共方法
 * @author      zhy    find404@foxmail.com
 * @createTime  2022年1月16日 13:30:00
 * @version     0.1.0 版本号
 */
import (
	"fmt"
	"os"
	"path/filepath"
)

// 获取系统临时目录
func GetTmpPath() (tmpPath string) {
	envPath := os.Getenv(clTemporaryPath)
	if envPath == clEmptyString {
		envPath = os.Getenv(clTmporaryPath)
	}
	tmpPath, err := filepath.EvalSymlinks(envPath)
	if err != nil {
		fmt.Printf("GetTmpDir,err: %v", err)
	}
	return
}
