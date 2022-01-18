package librarys

/**
 * 配置文件结合临时数据集
 * @author      zhy    find404@foxmail.com
 * @createTime  2022年1月16日 13:30:00
 * @version     0.0.1 版本号
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
	result := ClEmpty
	fileRawLen := len(fileRaw)
	if fileRawLen == 0 {
		return result
	}
	if find := strings.Contains(fileRaw, ClEqual); find {
		return result
	}
	if fileRaw[0:1] == ClLeftInBracket && fileRaw[fileRawLen-1:] == ClRightInBracket {
		result = fileRaw[1 : fileRawLen-1]
	}
	return result
}

//根据文件路径加载配置文件
func (ccf *Config) LoadIni(rootPath string, iniPath string) {
	if iniPath == ClEmpty || rootPath == ClEmpty {
		return
	}
	var indexPrefix string
	file, err := os.Open(rootPath + iniPath)
	if err != nil {
		fmt.Printf("Cannot open text file,err: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == ClEmpty {
			continue
		}
		//注释不处理
		if line[0:1] == ClPoundKey {
			continue
		}
		topIndex := ccf.getTopIndex(line)
		if topIndex != ClEmpty {
			indexPrefix = topIndex
			continue
		}
		eqIndex := strings.Index(line, ClEqual)
		if eqIndex == -1 {
			eqIndex = 0
		}

		lineKey := indexPrefix + ClUnderline + line[:eqIndex]
		lineVal := line[eqIndex+1:]

		GetTempDataInstance().SetMapStringString(lineKey, lineVal)
	}

	GetTempDataInstance().SetMapStringString(ClRootPath, iniPath)
}