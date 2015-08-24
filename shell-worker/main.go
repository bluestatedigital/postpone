// Shell worker.

package main

import (
    "os"
    
    flags "github.com/jessevdk/go-flags"
    log "github.com/Sirupsen/logrus"
    
    "github.com/jrallison/go-workers"
    
    "github.com/bluestatedigital/postpone/common"
    "github.com/bluestatedigital/postpone/tasks/shell"
)

var version string = "undef"

type Options struct {
    *common.Options

    Queue       string `long:"queue"       description:"queue for jobs to process" required:"true"`
    Concurrency int    `long:"concurrency"                                         default:"1"`
}

func main() {
    var opts Options
    
    _, err := flags.Parse(&opts)
    if err != nil {
        os.Exit(1)
    }
    
    if logFp := common.InitLogging(opts.Debug, opts.LogFile); logFp != nil {
        defer logFp.Close()
    }
    
    log.Debug("hi there! (tickertape tickertape)")
    log.Infof("version: %s", version)
    
    common.ConfigureWorkers(opts.RedisHost, opts.RedisPort, opts.RedisDb)
    
    workers.Process(opts.Queue, shell.Shell, opts.Concurrency)
    
    workers.Run()
}
