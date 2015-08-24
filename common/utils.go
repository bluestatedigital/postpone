package common

import (
    log "github.com/Sirupsen/logrus"
)

// Error assertion.
func CheckError(msg string, err error) {
    if err != nil {
        log.Fatalf("%s: %+v", msg, err)
    }
}
