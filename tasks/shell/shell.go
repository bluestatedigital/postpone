package shell

import (
    "os"
    "os/exec"
    "syscall"
    log "github.com/Sirupsen/logrus"

    "github.com/jrallison/go-workers"
)

func Shell(msg *workers.Msg) {
    logger := log.WithFields(log.Fields{
        "class": msg.Get("class").MustString(),
        "queue": msg.Get("queue").MustString(),
        "jid":   msg.Jid(),
    })
    
    path := msg.Get("args").Get("path").MustString()
    args := msg.Get("args").Get("args").MustStringArray()
    env  := msg.Get("args").Get("env").MustStringArray()
    
    if path == "" {
        logger.Panic("no command specified")
    }
    
    logger.Infof("shell job: path %s, args: %v, env: %v", path, args, env)
    
    cmd := exec.Command(path, args...)
    cmd.Env = os.Environ()
    for _, e := range env {
        cmd.Env = append(cmd.Env, e)
    }
    
    logger.Debugf("cmd: %+v", cmd)
    
    output, err := cmd.CombinedOutput()
    
    if len(output) > 0 {
        logger.WithField("type", "output").Info(string(output))
    }
    
    if err != nil {
        switch t := err.(type) {
            case *exec.ExitError:
                waitStatus := err.(*exec.ExitError).Sys().(syscall.WaitStatus)
                logger.Panicf("command exited with status %d, signal %d (%s)", waitStatus.ExitStatus(), waitStatus.Signal(), waitStatus.Signal())
            
            // case *exec.Error:
            //     logger.Panicf("Error: %+v", err)
            
            default:
                logger.Panicf("%T: %+v", t, err)
        }
    }
}
