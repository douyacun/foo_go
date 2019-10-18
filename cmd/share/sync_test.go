package share

import (
	"testing"
)

func TestSyncVariable(t *testing.T) {
	for i := 1; i < 10; i++ {
		syncVariable()
	}
}
