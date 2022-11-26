package log

import (
	"testing"
)

func TestInitLogger(t *testing.T) {
	var (
		path   = "./"
		level  = 4
		color  = true
		caller = true
	)
	InitLogger(WithPath(path), WithLevel(level), WithCaller(caller), WithColor(color))

	Logger.Infof("Infof: %d", 2)
	Logger.Warnf("Warnf: %d", 3)
	Logger.Errorf("Errorf: %d", 4)
}
