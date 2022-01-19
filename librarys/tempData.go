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
	SetMapStringInt(name string, value int)
	SetMapIntString(name int, value string)
	SetMapStringInterface(name string, value interface{})
	GetMapStringString(name string) string
	GetMapStringInt(name string) int
	GetMapIntString(name int) string
	GetMapStringInterface(name string) interface{}
}

//临时数据集
type tempData struct {
	mapStringString    map[string]string
	mapStringInt       map[string]int
	mapIntString       map[int]string
	mapStringInterface map[string]interface{}
	sync.Mutex
}

//初始化
func (td *tempData) initMap() {
	td.mapStringString = make(map[string]string)
	td.mapStringInt = make(map[string]int)
	td.mapIntString = make(map[int]string)
	td.mapStringInterface = make(map[string]interface{})
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

//设置string与int的map
func (td *tempData) SetMapStringInt(name string, value int) {
	if name == ClEmpty {
		return
	}
	td.Lock()
	defer td.Unlock()
	td.mapStringInt[name] = value
}

//设置int与string的map
func (td *tempData) SetMapIntString(name int, value string) {
	td.Lock()
	defer td.Unlock()
	td.mapIntString[name] = value
}

//设置string与Interface的map
func (td *tempData) SetMapStringInterface(name string, value interface{}) {
	if name == ClEmpty {
		return
	}
	td.Lock()
	defer td.Unlock()
	td.mapStringInterface[name] = value
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

//获取string与int的map
func (td *tempData) GetMapStringInt(name string) int {
	if name == ClEmpty {
		return 0
	}
	td.Lock()
	defer td.Unlock()
	return td.mapStringInt[name]
}

//获取int与string的map
func (td *tempData) GetMapIntString(name int) string {
	td.Lock()
	defer td.Unlock()
	return td.mapIntString[name]
}

//获取string与interface的map
func (td *tempData) GetMapStringInterface(name string) interface{} {
	if name == ClEmpty {
		return nil
	}
	td.Lock()
	defer td.Unlock()
	return td.mapStringInterface[name]
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
