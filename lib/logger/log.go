package logger

import (
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

// GlobalLogger is the global logger
var GlobalLogger = logrus.New()

// Fields is used by the customLogger object to output
// fields along with a message
type Fields map[string]interface{}

// Default options for the global logger
func init() {
	GlobalLogger.SetOutput(os.Stdout)
	GlobalLogger.SetLevel(logrus.DebugLevel)
}

// ErrorFields is a helper for logging fields to the global logger
func ErrorFields(message string, fields Fields) {
	GlobalLogger.WithFields(map[string]interface{}(fields)).Error(message)
}

// SetLogLevel sets the log level to the given level
func SetLogLevel(level string) {
	switch strings.ToLower(level) {
	case "info":
		GlobalLogger.SetLevel(logrus.InfoLevel)
	case "debug":
		GlobalLogger.SetLevel(logrus.DebugLevel)
	case "warn":
		GlobalLogger.SetLevel(logrus.WarnLevel)
	case "error":
		GlobalLogger.SetLevel(logrus.ErrorLevel)
	case "fatal":
		GlobalLogger.SetLevel(logrus.FatalLevel)
	case "panic":
		GlobalLogger.SetLevel(logrus.PanicLevel)
	default:
		GlobalLogger.SetLevel(logrus.DebugLevel)
		GlobalLogger.Warnf("Log level '%s' not recognised. Setting to Debug.", level)
	}
}

func PrintStack() string{
	var pcs  = make([]uintptr,240)
	n:=runtime.Callers(2,pcs)
	str:="\nstack:\n"
	for i :=0;i<n;i++{
		Func:=runtime.FuncForPC(pcs[i])
		file,line:=Func.FileLine(pcs[i])

		str+="\t"+file +" "+Func.Name()+", line "+ strconv.Itoa(line)+"\n"
	}
	return str
}
