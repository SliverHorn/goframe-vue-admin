package internal

import (
	"context"
	"fmt"
	"gf-vue-admin/server/library/global"
	"github.com/gogf/gf/frame/g"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"io/ioutil"
	"log"
	"os"
	"time"
)

// writer log writer interface
type writer interface {
	Printf(string, ...interface{})
}

type config struct {
	SlowThreshold time.Duration
	Colorful      bool
	LogLevel      logger.LogLevel
}

var (
	Discard = New(log.New(ioutil.Discard, "", log.LstdFlags), config{})
	Default = New(log.New(os.Stdout, "\r\n", log.LstdFlags), config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      logger.Warn,
		Colorful:      true,
	})
	Recorder = traceRecorder{Interface: Default, BeginAt: time.Now()}
)

func New(writer writer, config config) logger.Interface {
	var (
		infoStr      = "%s\n[info] "
		warnStr      = "%s\n[warn] "
		errStr       = "%s\n[error] "
		traceStr     = "%s\n[%.3fms] [rows:%v] %s"
		traceWarnStr = "%s %s\n[%.3fms] [rows:%v] %s"
		traceErrStr  = "%s %s\n[%.3fms] [rows:%v] %s"
	)

	if config.Colorful {
		infoStr = logger.Green + "%s\n" + logger.Reset + logger.Green + "[info] " + logger.Reset
		warnStr = logger.BlueBold + "%s\n" + logger.Reset + logger.Magenta + "[warn] " + logger.Reset
		errStr = logger.Magenta + "%s\n" + logger.Reset + logger.Red + "[error] " + logger.Reset
		traceStr = logger.Green + "%s\n" + logger.Reset + logger.Yellow + "[%.3fms] " + logger.BlueBold + "[rows:%v]" + logger.Reset + " %s"
		traceWarnStr = logger.Green + "%s " + logger.Yellow + "%s\n" + logger.Reset + logger.RedBold + "[%.3fms] " + logger.Yellow + "[rows:%v]" + logger.Magenta + " %s" + logger.Reset
		traceErrStr = logger.RedBold + "%s " + logger.MagentaBold + "%s\n" + logger.Reset + logger.Yellow + "[%.3fms] " + logger.BlueBold + "[rows:%v]" + logger.Reset + " %s"
	}

	return &customLogger{
		writer:       writer,
		config:       config,
		infoStr:      infoStr,
		warnStr:      warnStr,
		errStr:       errStr,
		traceStr:     traceStr,
		traceWarnStr: traceWarnStr,
		traceErrStr:  traceErrStr,
	}
}

type customLogger struct {
	writer
	config
	infoStr, warnStr, errStr            string
	traceStr, traceErrStr, traceWarnStr string
}

// LogMode log mode
func (c *customLogger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *c
	newLogger.LogLevel = level
	return &newLogger
}

// Info print info
func (c *customLogger) Info(ctx context.Context, message string, data ...interface{}) {
	if c.LogLevel >= logger.Info {
		c.Printf(c.infoStr+message, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Warn print warn messages
func (c *customLogger) Warn(ctx context.Context, message string, data ...interface{}) {
	if c.LogLevel >= logger.Warn {
		c.Printf(c.warnStr+message, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Error print error messages
func (c *customLogger) Error(ctx context.Context, message string, data ...interface{}) {
	if c.LogLevel >= logger.Error {
		c.Printf(c.errStr+message, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Trace print sql message
func (c *customLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if c.LogLevel > 0 {
		elapsed := time.Since(begin)
		switch {
		case err != nil && c.LogLevel >= logger.Error:
			sql, rows := fc()
			if rows == -1 {
				c.Printf(c.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				c.Printf(c.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		case elapsed > c.SlowThreshold && c.SlowThreshold != 0 && c.LogLevel >= logger.Warn:
			sql, rows := fc()
			slowLog := fmt.Sprintf("SLOW SQL >= %v", c.SlowThreshold)
			if rows == -1 {
				c.Printf(c.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				c.Printf(c.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		case c.LogLevel >= logger.Info:
			sql, rows := fc()
			if rows == -1 {
				c.Printf(c.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				c.Printf(c.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		}
	}
}

func (c *customLogger) Printf(message string, data ...interface{}) {
	if global.Config.Mysql.LogZap != "" {
		switch len(data) {
		case 0:
			g.Log().Info(message)
		case 1:
			g.Log().Info("gorm", g.Map{"src": data[0]})
		case 2:
			g.Log().Info("gorm", g.Map{"src": data[0]}, g.Map{"duration": data[1]})
		case 3:
			g.Log().Info("gorm", g.Map{"src": data[0]}, g.Map{"duration": data[1]}, g.Map{"rows": data[2]})
		case 4:
			g.Log().Info("gorm", g.Map{"src": data[0]}, g.Map{"duration": data[1]}, g.Map{"rows": data[2]}, g.Map{"sql": data[3]})
		}
		return
	}
	switch len(data) {
	case 0:
		c.writer.Printf(message, "")
	case 1:
		c.writer.Printf(message, data[0])
	case 2:
		c.writer.Printf(message, data[0], data[1])
	case 3:
		c.writer.Printf(message, data[0], data[1], data[2])
	case 4:
		c.writer.Printf(message, data[0], data[1], data[2], data[3])
	case 5:
		c.writer.Printf(message, data[0], data[1], data[2], data[3], data[4])
	}
}

type traceRecorder struct {
	logger.Interface
	BeginAt      time.Time
	SQL          string
	RowsAffected int64
	Err          error
}

func (t traceRecorder) New() *traceRecorder {
	return &traceRecorder{Interface: t.Interface, BeginAt: time.Now()}
}

func (t *traceRecorder) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	t.BeginAt = begin
	t.SQL, t.RowsAffected = fc()
	t.Err = err
}

func GenerateConfig() *gorm.Config {
	var _config = &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	switch global.Config.Mysql.LogZap {
	case "silent", "Silent":
		_config.Logger = Default.LogMode(logger.Silent)
	case "error", "Error":
		_config.Logger = Default.LogMode(logger.Error)
	case "warn", "Warn":
		_config.Logger = Default.LogMode(logger.Warn)
	case "info", "Info":
		_config.Logger = Default.LogMode(logger.Info)
	default:
		if global.Config.Mysql.LogMode {
			_config.Logger = Default.LogMode(logger.Info)
			break
		}
		_config.Logger = Default.LogMode(logger.Silent)
	}
	return _config
}
