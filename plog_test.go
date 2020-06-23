package plog_test

import (
    "flag"
    "testing"
    "github.com/practo/plog"
)

func init() {
    plog.InitFlags(nil)
}

func TestSeverityLogging(t *testing.T) {
    plog.Info("info")
}

func TestVerbosityLogging(t *testing.T) {
    flag.Set("v", "3")
    flag.Parse()
    plog.V(1).Info("verbosity=1")
    plog.V(2).Info("verbosity=2")
    plog.V(3).Info("verbosity=3")
}
