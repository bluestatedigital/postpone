package common

import (
    "os"
    "fmt"
    "syscall"
    
    log "github.com/Sirupsen/logrus"
    "github.com/jrallison/go-workers"
)

// Initializes logrus.  Returns a pointer to `os.File` that should be `Close`d
// via `defer` in the main func.
func InitLogging(debug bool, logFile string) *os.File {
    var logFp *os.File
    
    if debug {
        log.SetLevel(log.DebugLevel)
    }
    
    if logFile != "" {
        logFp, err := os.OpenFile(logFile, os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0600)
        CheckError(fmt.Sprintf("error opening %s", logFile), err)
        
        // ensure panic output goes to log file
        syscall.Dup2(int(logFp.Fd()), 1)
        syscall.Dup2(int(logFp.Fd()), 2)
        
        // log as JSON
        log.SetFormatter(&log.JSONFormatter{})
        
        // send output to file
        log.SetOutput(logFp)
    }
    
    workers.Logger = log.WithField("name", "worker")

    return logFp
}
