package common

import (
    "fmt"
    "github.com/jrallison/go-workers"
)

// Configures go-workers to connect to a particular Redis instance.
func ConfigureWorkers(host string, port int, db int) {
    workers.Configure(map[string]string{
        // location of redis instance
        "server":  fmt.Sprintf("%s:%d", host, port),

        // instance of the database
        "database":  fmt.Sprintf("%d", db),

        // number of connections to keep open with redis
        "pool":    "30",

        // unique process id for this instance of workers (for proper recovery of inprogress jobs on crash)
        "process": "1",
    })
}
