package log

import (
	"sync"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

const (
	logSuffix     = ".log"
	logTimeFormat = "2006-01-02 15:04:05.000"
)

var (
	Logger *logrus.Logger
	once   sync.Once
)

type opt struct {
	path       string
	level      int
	hasColor   bool
	hasCaller  bool
	timeFormat string
	suffix     string
	maxAge     time.Duration
}

func newOpt(opts ...Option) *opt {
	o := &opt{
		path:       "./",
		level:      int(logrus.TraceLevel),
		hasColor:   true,
		hasCaller:  true,
		timeFormat: logTimeFormat,
		suffix:     logSuffix,
		maxAge:     7 * 24 * time.Hour,
	}
	for _, fn := range opts {
		fn(o)
	}
	return o
}

type Option func(*opt)

func WithPath(path string) Option            { return func(o *opt) { o.path = path } }
func WithLevel(level int) Option             { return func(o *opt) { o.level = level } }
func WithColor(hasColor bool) Option         { return func(o *opt) { o.hasColor = hasColor } }
func WithCaller(hasCallor bool) Option       { return func(o *opt) { o.hasCaller = hasCallor } }
func WithTimeFormat(format string) Option    { return func(o *opt) { o.timeFormat = format } }
func WithSuffix(suffix string) Option        { return func(o *opt) { o.suffix = suffix } }
func WithMaxAge(maxAge time.Duration) Option { return func(o *opt) { o.maxAge = maxAge } }

// 配置日志行为
func InitLogger(opts ...Option) {
	o := newOpt(opts...)
	once.Do(func() {
		Logger.SetLevel(logrus.AllLevels[o.level%len(logrus.AllLevels)])
		Logger.SetReportCaller(o.hasCaller)
		Logger.SetFormatter(&logrus.TextFormatter{
			ForceColors:     o.hasColor,
			DisableColors:   !o.hasColor,
			TimestampFormat: o.timeFormat,
		})
		Logger.AddHook(fileLoggerHook(o))
	})
}

// 文件日志
func fileLoggerHook(o *opt) logrus.Hook {
	infoWriter, _ := rotatelogs.New(o.path+"%Y%m%d.info"+o.suffix,
		rotatelogs.WithMaxAge(o.maxAge),
		rotatelogs.WithRotationTime(24*time.Hour))
	warnWriter, _ := rotatelogs.New(o.path+"%Y%m%d.warn"+o.suffix,
		rotatelogs.WithMaxAge(o.maxAge),
		rotatelogs.WithRotationTime(24*time.Hour))
	errWriter, _ := rotatelogs.New(o.path+"%Y%m%d.error"+o.suffix,
		rotatelogs.WithMaxAge(o.maxAge),
		rotatelogs.WithRotationTime(24*time.Hour))

	fileFormatter := &logrus.TextFormatter{
		ForceColors:     false,
		DisableColors:   true,
		TimestampFormat: o.timeFormat,
	}

	return lfshook.NewHook(lfshook.WriterMap{
		logrus.InfoLevel:  infoWriter,
		logrus.WarnLevel:  warnWriter,
		logrus.ErrorLevel: errWriter,
	}, fileFormatter)
}

func init() {
	Logger = logrus.StandardLogger()
}
