package relimit

import (
	"fmt"
	"os/user"
	"strings"

	"github.com/realjf/cgroup"
	"github.com/realjf/utils"
)

type IReLimit interface {
	Start(cmdl string, args ...string) (output []byte, err error)
	Close()
	SetDebug(debug bool)
	GetCmd() *utils.Command
	GetCgroup() cgroup.ICgroup
	StartByPid(pid int) error
	SetUsername(username string) error
	SetUser(user *user.User) IReLimit
	SetNoSetGroups(noSetGroups bool)
}

type relimit struct {
	cg    cgroup.ICgroup
	cmd   *utils.Command
	debug bool
}

func NewRelimit(maxCpuUsage cgroup.Percent, maxMemoryInBytes cgroup.Memory, disableOOMKill bool) (IReLimit, error) {
	re := &relimit{
		debug: false,
		cmd:   utils.NewCmd(),
	}

	var err error
	err = re.cmd.CheckRoot()
	if err != nil {
		return nil, err
	}

	if cgroup.IsCgroupV2() {
		re.cg, err = cgroup.NewCgroup(cgroup.V2, cgroup.WithSlice("/"), cgroup.WithGroup(RandomGroup()))
	} else {
		re.cg, err = cgroup.NewCgroup(cgroup.V1, cgroup.WithName(RandomName()))
	}

	if err != nil {
		return nil, err
	}

	re.cg.SetOptions(cgroup.WithCPULimit(maxCpuUsage), cgroup.WithMemoryLimit(maxMemoryInBytes))
	if disableOOMKill {
		re.cg.SetOptions(cgroup.WithDisableOOMKiller())
	}

	err = re.cg.Create()
	if err != nil {
		return nil, err
	}

	return re, nil
}

func MustNewRelimit(maxCpuUsage cgroup.Percent, maxMemoryInBytes cgroup.Memory, disableOOMKill bool) IReLimit {
	re := &relimit{
		debug: false,
		cmd:   utils.NewCmd(),
	}

	var err error
	err = re.cmd.CheckRoot()
	if err != nil {
		panic(err)
	}

	if cgroup.IsCgroupV2() {
		re.cg, err = cgroup.NewCgroup(cgroup.V2, cgroup.WithSlice("/"), cgroup.WithGroup(RandomGroup()))
	} else {
		re.cg, err = cgroup.NewCgroup(cgroup.V1, cgroup.WithName(RandomName()))
	}
	if err != nil {
		panic(err)
	}

	re.cg.SetOptions(cgroup.WithCPULimit(maxCpuUsage), cgroup.WithMemoryLimit(maxMemoryInBytes))
	if disableOOMKill {
		re.cg.SetOptions(cgroup.WithDisableOOMKiller())
	}

	err = re.cg.Create()

	if err != nil {
		panic(err)
	}

	return re
}

func (r *relimit) SetDebug(debug bool) {
	r.debug = debug
}

func (r *relimit) GetCgroup() cgroup.ICgroup {
	return r.cg
}

func (r *relimit) GetCmd() *utils.Command {
	return r.cmd
}

func (r *relimit) Close() {
	defer r.cmd.Close()
	defer func() {
		err := r.cg.Close()
		if err != nil {
			if r.debug {
				fmt.Println(err)
			}

			return
		}
	}()
	if r.debug {
		fmt.Println("done!!!")
	}

}

func (r *relimit) StartByPid(pid int) error {
	err := r.cg.LimitPid(pid)
	if err != nil {
		if r.debug {
			fmt.Println(err)
		}
		return err
	}
	return nil
}

func (r *relimit) SetUser(user *user.User) IReLimit {
	r.cmd.SetUser(user)
	return r
}

func (r *relimit) SetUsername(username string) error {
	User, err := user.Lookup(username)
	if err != nil {
		return err
	}
	r.cmd.SetUser(User)
	return nil
}

func (r *relimit) SetNoSetGroups(noSetGroups bool) {
	r.cmd.SetNoSetGroups(noSetGroups)
}

func (r *relimit) Start(cmdl string, args ...string) (output []byte, err error) {

	_, err = r.cmd.Command(cmdl, args...)
	if err != nil {
		return nil, err
	}
	if r.debug {
		fmt.Printf("limit pid: %d\n", r.cmd.GetPid())
	}

	err = r.cg.LimitPid(r.cmd.GetPid())
	if err != nil {
		if r.debug {
			fmt.Println(err)
		}
		return nil, err
	}
	if r.debug {
		fmt.Printf("start run: %s %s\n", cmdl, strings.Join(args, " "))
	}

	output, err = r.cmd.Run()
	if err != nil {
		if r.debug {
			fmt.Println(err)
		}
	}
	return output, err
}
