package cgroup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCgroup(t *testing.T) {
	cases := map[string]struct {
		name    string
		slice   string
		group   string
		version CgroupVersion
	}{
		"v1": {
			name:    "test",
			version: V1,
		},
		"v2": {
			version: V2,
			slice:   "/",
			group:   "hello.slice",
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			cg, err := NewCgroup(tc.version, WithName(tc.name), WithSlice(tc.slice), WithGroup(tc.group))
			assert.NoError(t, err)
			if assert.NotNil(t, cg) {
				err = cg.Create()
				if assert.NoError(t, err) {
					err = cg.Close()
					assert.NoError(t, err)
				}

			}

		})
	}
}
