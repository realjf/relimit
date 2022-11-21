package relimit

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/process"
)

const (
	DefaultCPUUsage float64 = 80.0
)

type CPULimiter struct {
	// MaxCPUUsage is the CPU usage threshold, in percent, above which
	// the CPU is considered as "hot". This value can be set while the
	// measurement process is running. Default: 80.0
	MaxCPUUsage float64
}

func (c *CPULimiter) Start() {
	if c.MaxCPUUsage == 0.0 {
		c.MaxCPUUsage = DefaultCPUUsage
	}

	go c.run()
}

func (c *CPULimiter) run() {

}

func (c *CPULimiter) Stop() {

}

func getGlobalCpuTimes() cpu.TimesStat {
	ts, err := cpu.Times(false)
	if err != nil {
		panic(err)
	}
	return ts[0]
}

func getProcessCpuTimes(proc *process.Process) cpu.TimesStat {
	t, err := proc.Times()
	if err != nil {
		panic(err)
	}
	return *t
}

func busyFromTimes(t cpu.TimesStat) (busy, all float64) {
	busy = t.User + t.System + t.Nice + t.Iowait + t.Irq + t.Softirq + t.Steal + t.Guest + t.GuestNice
	all = busy + t.Idle
	return
}

func getCPUUsage(busy1, all1, busy2, all2 float64) float64 {
	if all1 == all2 {
		return 0.0
	}
	usage := ((busy2 - busy1) / (all2 - all1)) * 100.0
	return usage
}
