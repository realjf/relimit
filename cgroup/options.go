package cgroup

type Option func(cgroup ICgroup)

// for cgroup v1
func WithName(name string) Option {
	return func(cgroup ICgroup) {
		if c, ok := cgroup.Instance().(*cgroupImplV1); ok {
			c.name = name
		}
	}
}

// for cgroup v2
func WithSlice(slice string) Option {
	return func(cgroup ICgroup) {
		if c, ok := cgroup.Instance().(*cgroupImplV2); ok {
			c.slice = slice
		}
	}
}

// for cgroup v2
func WithGroup(group string) Option {
	return func(cgroup ICgroup) {
		if c, ok := cgroup.Instance().(*cgroupImplV2); ok {
			c.group = group
		}
	}
}
