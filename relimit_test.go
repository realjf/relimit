package relimit

import (
	"fmt"
	"os"
	"testing"

	"github.com/realjf/cgroup"
)

func TestRelimit(t *testing.T) {
	re := MustNewRelimit(80, 1000*cgroup.Megabyte, true)
	re.SetDebug(true)
	re.GetCmd().SetDebug(true)
	defer re.Close()
	re.SetUsername(os.Getenv("SUDO_USER"))
	re.SetNoSetGroups(true)
	args := []string{"/home/realjf/Downloads/5f8eb954-fa7e-4f41-a935-22c6892865be.epub", "/home/realjf/Downloads/5f8eb954-fa7e-4f41-a935-22c6892865be.pdf"}
	out, err := re.Start("ebook-convert", args...)
	// args := []string{"--cpu", "1", "--vm", "1", "--vm-bytes", "20M", "--timeout", "10s", "--vm-keep"}
	// out, err := re.Start("stress", args...)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("output: %s\n", out)
	}

}
