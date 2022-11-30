package cgroup

import "errors"

type CgroupVersion string

const (
	V1 CgroupVersion = "v1"
	V2 CgroupVersion = "v2"
)

// String is used both by fmt.Print and by Cobra in help textv
func (e *CgroupVersion) String() string {
	return string(*e)
}

// Set must have pointer receiver so it doesn't change the value of a copy
func (e *CgroupVersion) Set(v string) error {
	switch v {
	case "v1":
		*e = CgroupVersion(v)
		return nil
	case "v2":
		*e = CgroupVersion(v)
		return nil
	default:
		return errors.New(`must be one of ["v1", "v2"]`)
	}
}

// Type is only used in help text
func (e *CgroupVersion) Type() string {
	return "cgroup version"
}
