package utils

import (
	"gf-vue-admin/server/app/api/response"
	"gf-vue-admin/server/library/constant"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"runtime"
	"time"
)

var Server = new(server)

type server struct {
	err     error
	_server response.Server
}

func (s *server) Data() (*response.Server, error) {
	// Os
	s._server.Os.GOOS = runtime.GOOS
	s._server.Os.NumCPU = runtime.NumCPU()
	s._server.Os.Compiler = runtime.Compiler
	s._server.Os.GoVersion = runtime.Version()
	s._server.Os.NumGoroutine = runtime.NumGoroutine()

	// CPU
	if cores, err := cpu.Counts(false); err != nil {
		return &s._server, err
	} else {
		s._server.Cpu.Cores = cores
	}
	if cpus, err := cpu.Percent(time.Duration(200)*time.Millisecond, true); err != nil {
		return &s._server, err
	} else {
		s._server.Cpu.Cpus = cpus
	}

	// 内存
	if _mem, err := mem.VirtualMemory(); err != nil {
		return &s._server, err
	} else {
		s._server.Rrm.UsedMB = int(_mem.Used) / constant.MB
		s._server.Rrm.TotalMB = int(_mem.Total) / constant.MB
		s._server.Rrm.UsedPercent = int(_mem.UsedPercent)
	}

	// 磁盘空间
	if _dist, err := disk.Usage("/"); err != nil {
		return &s._server, err
	} else {
		s._server.Disk.UsedMB = int(_dist.Used) / constant.MB
		s._server.Disk.UsedGB = int(_dist.Used) / constant.GB
		s._server.Disk.TotalMB = int(_dist.Total) / constant.MB
		s._server.Disk.TotalGB = int(_dist.Total) / constant.GB
		s._server.Disk.UsedPercent = int(_dist.UsedPercent)
	}
	return &s._server, nil
}
