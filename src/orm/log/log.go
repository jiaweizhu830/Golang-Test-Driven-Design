package log

import (
	"io/ioutil"
	"log"
	"os"
	"sync"
)

// declaring multiple vars at a time
var (
	// LstdFlags     = Ldate | Ltime // initial values for the standard logger
	// Lshortfile: final file name element and line number: d.go:23. overrides Llongfile
	errorLog = log.New(os.Stdout, "\033[31m[error]\033[0m ", log.LstdFlags|log.Lshortfile)
	infoLog  = log.New(os.Stdout, "\033[34m[info]\033[0m ", log.LstdFlags|log.Lshortfile)
	loggers  = []*log.Logger{errorLog, infoLog}
	mu       sync.Mutex
)

// log methods
var (
	Error  = errorLog.Println
	Errorf = errorLog.Printf
	Info   = infoLog.Println
	Infof  = infoLog.Printf
)

// log levels
const (
	// iota enum
	InfoLevel  = iota // 0
	ErrorLevel        // 1
	Disabled          // 2
)

// SetLevel controls log level
func SetLevel(level int) {
	mu.Lock()
	defer mu.Unlock()

	for _, logger := range loggers {
		logger.SetOutput(os.Stdout)
	}

	if ErrorLevel < level {
		// not log
		errorLog.SetOutput(ioutil.Discard)
	}
	if InfoLevel < level {
		// not log
		infoLog.SetOutput(ioutil.Discard)
	}

}
