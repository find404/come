package librarys

/**
 * 命令执行包
 * @author      zhy    find404@foxmail.com
 * @createTime  2022年1月19日 22:01:00
 * @version     0.1.0 版本号
 */
import (
	"os/exec"
	"sync"
)

//对外接口
type CmdExecSingleton interface {
	ExecGitCmd(arg ...string) (string, error)
	SetExecDir(name string)
}

//命令行执行类
type cmdExec struct {
	execDir string
	sync.Mutex
}

//执行GIT命令行
func (ce *cmdExec) ExecGitCmd(arg ...string) (string, error) {
	ce.Lock()
	defer ce.Unlock()

	cmd := exec.Command("git", arg...)
	cmd.Dir = ce.execDir
	msg, err := cmd.CombinedOutput()
	cmd.Run()
	return string(msg), err
}

//设置GIT命令行执行目录
func (ce *cmdExec) SetExecDir(name string) {
	if name == ClEmpty {
		return
	}
	ce.Lock()
	defer ce.Unlock()
	ce.execDir = name
}

var (
	cmdExecInstance *cmdExec
	cmdExecOnce     sync.Once
)

//获取实例
func GetCmdExecInstance() CmdExecSingleton {
	cmdExecOnce.Do(func() {
		cmdExecInstance = &cmdExec{}
	})
	return cmdExecInstance
}
