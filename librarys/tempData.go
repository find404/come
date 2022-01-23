package librarys

/**
 * 当前包临时储存
 * @author      zhy    find404@foxmail.com
 * @createTime  2022年1月16日 13:30:00
 * @version     0.1.0 版本号
 */
import (
	"sync"
)

//对外接口
type TempDataSingleton interface {
	SetMapStringString(name string, value string)
	SetMapStringKeyString(pName string, name string, value string)
	SetMapStringMapString(name string, value string)
	GetMapStringString(name string) string
	GetMapStringKeyString(pName string, name string) string
	GetMapStringMapString(name string) []string
}

//临时数据集
type tempData struct {
	mapStringString    map[string]string
	mapStringKeyString map[string]map[string]string
	mapStringMapString map[string][]string
	sync.Mutex
}

//初始化
func (td *tempData) initMap() {
	td.mapStringString = make(map[string]string)
	td.mapStringKeyString = make(map[string]map[string]string)
	td.mapStringMapString = make(map[string][]string)
}

//设置二维键值对的map，存储格式如下
//{
//  "键值":{ "下级键" :"值" }
//}
func (td *tempData) SetMapStringKeyString(pName string, name string, value string) {
	if pName == ClEmpty {
		return
	}
	td.Lock()
	defer td.Unlock()

	//设置值为空，为删除
	if name == ClEmpty {
		delete(td.mapStringKeyString, pName)
		return
	}
	if _, ok := td.mapStringKeyString[pName]; !ok {
		td.mapStringKeyString[pName] = make(map[string]string)
	}
	td.mapStringKeyString[pName][name] = value
}

//设置一键一map的map，存储格式如下
//{
//  "键值":[ "值" ]
//}
func (td *tempData) SetMapStringMapString(name string, value string) {
	if name == ClEmpty {
		return
	}
	td.Lock()
	defer td.Unlock()

	//设置值为空，为删除
	if value == ClEmpty {
		delete(td.mapStringMapString, name)
		return
	}
	td.mapStringMapString[name] = append(td.mapStringMapString[name], value)
}

//设置string与string的map
func (td *tempData) SetMapStringString(name string, value string) {
	if name == ClEmpty {
		return
	}
	td.Lock()
	defer td.Unlock()

	td.mapStringString[name] = value
}

//获取一键一map的map
func (td *tempData) GetMapStringMapString(name string) []string {
	if name == ClEmpty {
		return []string{}
	}
	td.Lock()
	defer td.Unlock()

	return td.mapStringMapString[name]
}

//获取二维键值对的map
func (td *tempData) GetMapStringKeyString(pName string, name string) string {
	if pName == ClEmpty {
		return ClEmpty
	}
	td.Lock()
	defer td.Unlock()

	return td.mapStringKeyString[pName][name]
}

//获取string与string的map
func (td *tempData) GetMapStringString(name string) string {
	if name == ClEmpty {
		return ClEmpty
	}
	td.Lock()
	defer td.Unlock()

	return td.mapStringString[name]
}

var (
	tempDataInstance *tempData
	once             sync.Once
)

//获取实例
func GetTempDataInstance() TempDataSingleton {
	once.Do(func() {
		tempDataInstance = &tempData{}
		tempDataInstance.initMap()
	})
	return tempDataInstance
}
