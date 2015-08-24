package shell_test

import (
    . "github.com/bluestatedigital/postpone/tasks/shell"

    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    
    // "github.com/stretchr/testify/mock"
    
    "github.com/jrallison/go-workers"
)

var _ = Describe("Submodule", func() {
    It("runs a command", func() {
        msg, err := workers.NewMsg(`{
            "jid": "e13897e4829c647bbcf7c6c2",
            "queue": "some-queue",
            "class": "shell or something",
            "enqueued_at": 1440303319.004147,
            "at": 1440303319.0041466,
            "args": {
                "path": "/bin/sh",
                "args": [
                    "-c",
                    "env"
                ],
                "env": [
                    "FOO=bar"
                ]
            }
        }`)
        
        Expect(err).To(BeNil())
        
        Shell(msg)
    })

    It("runs another command", func() {
        msg, err := workers.NewMsg(`{
            "jid": "e13897e4829c647bbcf7c6c2",
            "queue": "some-queue",
            "class": "shell or something",
            "enqueued_at": 1440303319.004147,
            "at": 1440303319.0041466,
            "args": {
                "path": "env"
            }
        }`)
        
        Expect(err).To(BeNil())
        
        Shell(msg)
    })
})
