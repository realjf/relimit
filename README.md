# relimit
Relimit is a tool that limits the CPU and memory usage of a process.(cpu和内存使用限制工具)


### Usage
```go
import (
 "fmt"
 "github.com/realjf/cgroup"
)

func main() {
 re := MustNewRelimit(30, 8*cgroup.Megabyte, true)
 defer re.Close()
 args := []string{"--cpu", "1", "--vm", "1", "--vm-bytes", "20M", "--timeout", "10s", "--vm-keep"}
 out, err := re.Start("stress", args...)
 if err != nil {
   fmt.Println(err.Error())
 }
 fmt.Printf("output: %s\n", out)
}

```

> Please use it as root or sudo.
