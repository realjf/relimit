package relimit

import (
	"fmt"
	"testing"

	"github.com/realjf/cgroup"
)

func TestRelimit(t *testing.T) {
	re := MustNewRelimit(30, 8*cgroup.Megabyte, true)
	defer re.Close()
	args := []string{"--cpu", "1", "--vm", "1", "--vm-bytes", "20M", "--timeout", "10s", "--vm-keep"}
	out, err := re.Start("stress", args...)
	if err != nil {
		t.Fatalf(err.Error())
	}
	fmt.Printf("output: %s\n", out)
}
