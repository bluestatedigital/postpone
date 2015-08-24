package shell_test

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"

    "github.com/Sirupsen/logrus"

    "testing"
)

func TestShell(t *testing.T) {
    RegisterFailHandler(Fail)

    logrus.SetLevel(logrus.PanicLevel)

    RunSpecs(t, "Shell Suite")
}
