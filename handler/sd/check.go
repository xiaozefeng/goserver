package sd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"net/http"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

// HealthCheck shows `OK` as the ping-pong result.
func HealthCheck(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

// DiskCheck checks the disk usage
func DiskCheck(c *gin.Context) {
	usageStatus, _ := disk.Usage("/")

	used := int(usageStatus.Used)
	total := int(usageStatus.Total)

	usedMB := used / MB
	usedGB := used / GB
	totalMB := total / MB
	totalGB := total / GB
	usedPercent := int(usageStatus.UsedPercent)

	status := http.StatusOK
	text := "OK"

	if usedPercent >= 95 {
		status = http.StatusOK
		text = "CRITICAL"
	} else if usedPercent >= 90 {
		status = http.StatusTooManyRequests
		text = "warning"
	}

	message := fmt.Sprintf("%s - Free space: %dMB (%dGB) / %dMB (%dGB) | Used: %d%%",
		text,
		usedMB,
		usedGB,
		totalMB,
		totalGB,
		usedPercent)

	c.String(status, message)
}

// CPUCheck checks the cpu usage
func CPUCheck(c *gin.Context) {
	cores, _ := cpu.Counts(false)
	a, _ := load.Avg()
	l1 := a.Load1
	l5 := a.Load5
	l15 := a.Load15

	status := http.StatusOK
	text := "OK"

	if l5 >= float64(cores-1) {
		status = http.StatusInternalServerError
		text = "CRITICAL"
	} else if l5 >= float64(cores-2) {
		status = http.StatusTooManyRequests
		text = "WARNING"
	}

	message := fmt.Sprintf("%s - Load average: %.2f, %.2f, %.2f | Cores: %d",
		text,
		l1,
		l5,
		l15,
		cores)
	c.String(status, message)
}

// RAMCheck checks the disk usage
func RAMCheck(c *gin.Context) {
	vms, _ := mem.VirtualMemory()

	used := int(vms.Used)
	total := int(vms.Total)

	usedMB := used / MB
	usedGB := used / GB
	totalMB := total / MB
	totalGB := total / GB
	usedPercent := int(vms.UsedPercent)

	status := http.StatusOK
	text := "OK"

	if usedPercent >= 95 {
		status = http.StatusInternalServerError
		text = "CRITICAL"
	} else if usedPercent >= 90 {
		status = http.StatusTooManyRequests
		text = "WARNING"
	}

	message := fmt.Sprintf("%s - Free space: %dMB (%dGB) / %dMB (%dGB) | Used: %d%% ",
		text,
		usedMB,
		usedGB,
		totalMB,
		totalGB,
		usedPercent)
	c.String(status, message)
}
