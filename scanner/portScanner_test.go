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

func testSerialScanner(t *testing.T) {
	scanner := NewSerialPortScanner("127.0.0.1", &testPinger{})
	res := scanner.Scan()
	assert.ElementsMatch(t, []int{80, 8080}, res)
}

func testParallelPortScanner(t *testing.T) {
	scanner := NewParallelPortScanner("127.0.0.1", &testPinger{}, 5)
	res := scanner.Scan()
	assert.ElementsMatch(t, []int{80, 8080}, res)
}
