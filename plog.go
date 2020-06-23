package plog

import (
	"os"
	"fmt"
	"flag"
	"sync"
	"k8s.io/klog/v2"
)

type loggingT struct {
	mu sync.Mutex
	hooks Hooks
}

var logging loggingT

func init() {
	logging.hooks = Hooks{}
}

func AddHook(hook Hook) {
	logging.addHook(hook)
}

// klog does not expose this, hence the duplication
type severity int32

// klog does not expose this, hence the duplication
const (
	infoLog severity = iota
	warningLog
	errorLog
	fatalLog
	numSeverity = 4
)

// klog does not expose this, hence the duplication
var severityName = []string{
	infoLog:    "INFO",
	warningLog: "WARNING",
	errorLog:   "ERROR",
	fatalLog:   "FATAL",
}

// addHook adds a Hook to logging
func (l *loggingT) addHook(hook Hook) {
	l.mu.Lock()
	defer l.mu.Unlock()
	err := l.hooks.Add(hook)
	if err != nil {
		l.exit(err)
	}
}

func (l *loggingT) fireHooks(s severity) {
	err := l.hooks.Fire(s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to fire hook: %v\n", err)
	}
}

func (l *loggingT) exit(err error) {
	fmt.Fprintf(os.Stderr, "log: exiting because of error: %s\n", err)
	// If logExitFunc is set, we do that instead of exiting.
	if logExitFunc != nil {
		logExitFunc(err)
		return
	}
	l.flushAll()
	os.Exit(2)
}

func InitFlags(flagset *flag.FlagSet) {
	klog.InitFlags(flagset)
}

type Level klog.Level

type Verbose struct {
	enabled bool
}

func V(level klog.Level) Verbose {
	v := klog.V(level)
	return Verbose{
		enabled: v.Enabled(),
	}
}

func (v Verbose) Enabled() bool {
	return v.enabled
}

func (v Verbose) Info(args ...interface{}) {
	if v.enabled {
		klog.Info(args...)
	}
}

func (v Verbose) Infoln(args ...interface{}) {
	if v.enabled {
		klog.Infoln(args...)
	}
}

func (v Verbose) Infof(format string, args ...interface{}) {
	if v.enabled {
		klog.Infof(format, args...)
	}
}

func (v Verbose) InfoS(msg string, keysAndValues ...interface{}) {
	if v.enabled {
		klog.InfoS(msg, keysAndValues...)
	}
}

func (v Verbose) Error(err error, msg string, args ...interface{}) {
	if v.enabled {
		klog.ErrorS(err, msg, args...)
	}
}

func Info(args ...interface{}) {
	klog.Info(args...)
}

func InfoDepth(depth int, args ...interface{}) {
	klog.InfoDepth(depth, args...)
}

func Infoln(args ...interface{}) {
	klog.Infoln(args...)
}

func Infof(format string, args ...interface{}) {
	klog.Infof(format, args...)
}

func InfoS(msg string, keysAndValues ...interface{}) {
	klog.InfoS(msg, keysAndValues...)
}

func Warning(args ...interface{}) {
	klog.Warning(args...)
}

func WarningDepth(depth int, args ...interface{}) {
	klog.WarningDepth(depth, args...)
}

func Warningln(args ...interface{}) {
	klog.Warningln(args...)
}

func Warningf(format string, args ...interface{}) {
	klog.Warningf(format, args...)
}

func Error(args ...interface{}) {
	klog.Error(args...)
}

func ErrorDepth(depth int, args ...interface{}) {
	klog.ErrorDepth(depth, args...)
}

func Errorln(args ...interface{}) {
	klog.Errorln(args...)
}

func Errorf(format string, args ...interface{}) {
	klog.Errorf(format, args...)
}

func ErrorS(err error, msg string, keysAndValues ...interface{}) {
	klog.ErrorS(err, msg, keysAndValues...)
}

func Fatal(args ...interface{}) {
	klog.Fatal(args...)
}

func FatalDepth(depth int, args ...interface{}) {
	klog.FatalDepth(depth, args...)
}

func Fatalln(args ...interface{}) {
	klog.Fatalln(args...)
}

func Fatalf(format string, args ...interface{}) {
	klog.Fatalf(format, args...)
}

func Exit(args ...interface{}) {
	klog.Exit(args...)
}

func ExitDepth(depth int, args ...interface{}) {
	klog.ExitDepth(depth, args...)
}

func Exitln(args ...interface{}) {
	klog.Exitln(args...)
}

func Exitf(format string, args ...interface{}) {
	klog.Exitf(format, args...)
}
