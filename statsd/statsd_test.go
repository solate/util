package statsd

import "testing"

func TestStatsd_StatsdPush(t *testing.T) {
	s := NewStatsd("udp", "127.0.0.1:8125", 0)
	err := s.Init()
	if err != nil {
		t.Error(err.Error())
		return
	}
	s.StatsdPush("test", map[string]string{"foo": "bar"}, "core", "suffix")
}
