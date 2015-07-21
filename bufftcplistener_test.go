package buffstreams

import (
	"strconv"
	"testing"
)

func TestListenBuffTCPUsesDefaultMessageSize(t *testing.T) {
	cfg := BuffTCPListenerConfig{
		Address: FormatAddress("", strconv.Itoa(5031)),
	}
	buffM, err := ListenBuffTCP(cfg)
	if err != nil {
		t.Errorf("Could not Listen on port %d: %s", cfg.MaxMessageSize, err.Error())
	}
	if buffM.maxMessageSize != DefaultMaxMessageSize {
		t.Errorf("Expected Max Message Size to be %d, actually got %d", DefaultMaxMessageSize, buffM.maxMessageSize)
	}
}

func TestListenBuffTCPUsesSpecifiedMaxMessageSize(t *testing.T) {
	cfg := BuffTCPListenerConfig{
		MaxMessageSize: 8196,
		Address:        FormatAddress("", strconv.Itoa(5032)),
	}
	buffM, err := ListenBuffTCP(cfg)
	if err != nil {
		t.Errorf("Could not Listen on port %d: %s", cfg.MaxMessageSize, err.Error())
	}
	if buffM.maxMessageSize != cfg.MaxMessageSize {
		t.Errorf("Expected Max Message Size to be %d, actually got %d", cfg.MaxMessageSize, buffM.maxMessageSize)
	}
}