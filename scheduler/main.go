// Job scheduler process.

package main

import (
    "os"
    "io/ioutil"
    
    flags "github.com/jessevdk/go-flags"
    log "github.com/Sirupsen/logrus"
    
    "github.com/jrallison/go-workers"
    
    "github.com/bluestatedigital/postpone/common"
    
    "gopkg.in/robfig/cron.v2"
    "github.com/ghodss/yaml"
)

/*
Field name   | Mandatory? | Allowed values  | Allowed special characters
----------   | ---------- | --------------  | --------------------------
Seconds      | No         | 0-59            | * / , -
Minutes      | Yes        | 0-59            | * / , -
Hours        | Yes        | 0-23            | * / , -
Day of month | Yes        | 1-31            | * / , - ?
Month        | Yes        | 1-12 or JAN-DEC | * / , -
Day of week  | Yes        | 0-6 or SUN-SAT  | * / , - ?

Entry                  | Description                                | Equivalent To
-----                  | -----------                                | -------------
@yearly (or @annually) | Run once a year, midnight, Jan. 1st        | 0 0 0 1 1 *
@monthly               | Run once a month, midnight, first of month | 0 0 0 1 * *
@weekly                | Run once a week, midnight on Sunday        | 0 0 0 * * 0
@daily (or @midnight)  | Run once a day, midnight                   | 0 0 0 * * *
@hourly                | Run once an hour, beginning of hour        | 0 0 * * * *
*/

type Schedule struct {
    Spec  string
    Queue string
    Class string
    Args  interface{}
}

type Opts struct {
    *common.Options
    
    Config string `long:"config" description:"config file (yaml or json)" required:"true"`
}

var version string = "undef"

func main() {
    var opts Opts
    
    _, err := flags.Parse(&opts)
    if err != nil {
        os.Exit(1)
    }
    
    if logFp := common.InitLogging(opts.Debug, opts.LogFile); logFp != nil {
        defer logFp.Close()
    }
    
    log.Debug("hi there! (tickertape tickertape)")
    log.Infof("version: %s", version)
    
    cfgBytes, err := ioutil.ReadFile(opts.Config)
    common.CheckError("reading config file", err)
    
    var schedule []Schedule
    common.CheckError("parsing config file", yaml.Unmarshal(cfgBytes, &schedule))
    
    log.Errorf("%+v", schedule)
    
    common.ConfigureWorkers(opts.RedisHost, opts.RedisPort, opts.RedisDb)

    c := cron.New()
    
    for _, e := range schedule {
        // some weird scoping going on here
        func (entry Schedule) {
            c.AddFunc(entry.Spec, func() {
                jid, err := workers.Enqueue(entry.Queue, entry.Class, entry.Args)
                common.CheckError("scheduling something", err)
                
                log.WithField("something", entry.Spec).Infof("submitted job %s", jid)
            })
        }(e)
    }
    
    c.Start()
    
    select {}
}
