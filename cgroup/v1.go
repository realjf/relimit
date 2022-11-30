package cgroup

import (
	"errors"
	"fmt"

	"github.com/containerd/cgroups"
	"github.com/containerd/cgroups/v3/cgroup1"
	"github.com/opencontainers/runtime-spec/specs-go"
)

type CgroupV1 interface {
	ICgroup
}

type cgroupImplV1 struct {
	cg      cgroups.Cgroup
	name    string
	version CgroupVersion
}

func NewCgroupImplV1() *cgroupImplV1 {
	return &cgroupImplV1{
		version: V1,
	}
}

func (c *cgroupImplV1) Version() CgroupVersion {
	return c.version
}

func (c *cgroupImplV1) SetOptions(options ...Option) {
	for _, opt := range options {
		opt(c)
	}
}

func (c *cgroupImplV1) Instance() any {
	return c
}

func (c *cgroupImplV1) SetName(name string) {
	c.name = name
}

func (c *cgroupImplV1) Close() error {
	return c.cg.Delete()
}

func (c *cgroupImplV1) Load() error {
	cgroup1.Default()
	var err error
	c.cg, err = cgroups.Load(cgroups.V1, cgroups.StaticPath(fmt.Sprintf("/%s", c.name)))
	return err
}

func (c *cgroupImplV1) Create() error {
	if c.name == "" {
		return errors.New("name is empty")
	}
	var err error
	c.cg, err = cgroups.New(cgroups.V1, cgroups.StaticPath(fmt.Sprintf("/%s", c.name)), &specs.LinuxResources{
		CPU:            &specs.LinuxCPU{},
		Memory:         &specs.LinuxMemory{},
		BlockIO:        &specs.LinuxBlockIO{},
		Network:        &specs.LinuxNetwork{},
		Pids:           &specs.LinuxPids{},
		Devices:        make([]specs.LinuxDeviceCgroup, 0),
		Rdma:           make(map[string]specs.LinuxRdma),
		HugepageLimits: make([]specs.LinuxHugepageLimit, 0),
	})
	return err
}
