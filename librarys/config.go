package librarys

/**
 * 配置文件结合临时数据集
 * @author      zhy    find404@foxmail.com
 * @createTime  2022年1月16日 13:30:00
 * @version     0.1.0 版本号
 */
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//空配置
type Config struct {
}

//获取上级索引配置 eg : []
func (ccf *Config) getTopIndex(fileRaw string) string {
	result := clEmptyString
	fileRawLen := len(fileRaw)
	if fileRawLen == 0 {
		return result
	}
	if find := strings.Contains(fileRaw, clEqualsSymbol); find {
		return result
	}
	if fileRaw[0:1] == clInBracketsLeft && fileRaw[fileRawLen-1:] == clInBracketsRight {
		result = fileRaw[1 : fileRawLen-1]
	}
	return result
}

//根据文件路径加载配置文件
func (ccf *Config) LoadIni(iniPath string) {
	if iniPath == clEmptyString {
		return
	}
	var indexPrefix string
	file, err := os.Open(iniPath)
	if err != nil {
		fmt.Printf("Cannot open text file,err: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == clEmptyString {
			continue
		}
		//注释不处理
		if line[0:1] == clPoundKeySymbol {
			continue
		}
		topIndex := ccf.getTopIndex(line)
		if topIndex != clEmptyString {
			indexPrefix = topIndex
			continue
		}
		eqIndex := strings.Index(line, clEqualsSymbol)
		if eqIndex == -1 {
			eqIndex = 0
		}

		lineKey := indexPrefix + clUnderlineSymbol + line[:eqIndex]
		lineVal := line[eqIndex+1:]

		GetTempDataInstance().SetMapStringString(lineKey, lineVal)
	}

	GetTempDataInstance().SetMapStringString("rootPath", iniPath)
}

//手动设置值
func (ccf *Config) SetConfigString(name string, value string) {
	GetTempDataInstance().SetMapStringString(name, value)
}
