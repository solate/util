package logger

import "testing"

func TestNewLogger(t *testing.T) {

	logger, err := NewLogger("debug", "./out.log")
	if err != nil {
		t.Error(err)
		return
	}

	logger.Debug("test Debug")
	logger.Info("test Info")
	logger.Warn("test Warn")
	logger.Error("test Error")
	//logger.DPanic("test DPanic")
	//logger.Panic("test Panic")
	//logger.Fatal("test Fatal")

}
