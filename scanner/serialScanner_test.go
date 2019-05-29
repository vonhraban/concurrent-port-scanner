package scanner

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type testPinger struct{}

func (p *testPinger) Ping(ip string, port int) bool {
	if ip == "127.0.0.1" {
		switch port {
		case 80:
			return true

		case 8080:
			return true
		}
	}

	return false
}

func TestSerialScanner(t *testing.T) {
	scanner := NewSerialPortScanner("127.0.0.1", &testPinger{})
	res := scanner.Scan()
	assert.True(t, res[80])
	assert.True(t, res[8080])
	assert.False(t, res[21])
	assert.False(t, res[9000])
}
