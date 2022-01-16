# come/Librarys
come is come go is go, come frame框架下面，自用的librarys，组件库

## 依赖
* golang > 1.6

## 功能清单

config.go:配置文件相关
```
LoadIni(pathName string)
getTopIndex(fileRaw string)
SetConfigString(name string, value string)
```

consts.go:定义常量
funcs.go:内置方法集
```
GetTmpPath()
```
tempData.go 临时数据集
```
GetTempDataInstance
SetMapStringString(name string, value string)
SetMapStringInt(name string, value int)
SetMapIntString(name int, value string)
SetMapStringInterface(name string, value interface{})
GetMapStringString(name string) string
GetMapStringInt(name string) int
GetMapIntString(name int) string
GetMapStringInterface(name string) interface{}
```


