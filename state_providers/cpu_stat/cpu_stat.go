package cpu_stat

import (
	"github.com/mackerelio/go-osstat/cpu"
	"time"
)

func Get() (int, error) {
	before, err := cpu.Get()
	if err != nil {
		return 0, err
	}

	time.Sleep(time.Duration(1) * time.Second)

	after, err2 := cpu.Get()
	if err2 != nil {
		return 0, nil
	}

	total := float64(after.Total - before.Total)

	cpuUser := float64(after.User-before.User) / total * 100

	return int(cpuUser), nil
}
