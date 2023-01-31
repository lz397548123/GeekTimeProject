package mutex1

import (
	"testing"
)

func TestMutex(t *testing.T) {
	m := NewMutex()
	ok := m.TryLock()
	t.Logf("locked v %v\n", ok)
	ok = m.TryLock()
	t.Logf("locked %v\n", ok)
}
