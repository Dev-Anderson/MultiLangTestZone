package newlogger

import (
	"runtime"
	"strconv"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

func getMemory() string {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	valueMem := strconv.FormatUint(mem.Alloc, 10)

	return valueMem
}

func getCPU() (float64, error) {
	info, err := cpu.Percent(time.Second, false)
	if err != nil {
		return 0, err
	}

	return info[0], nil
}
