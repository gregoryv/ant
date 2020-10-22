package ant

import "testing"

func TestSetFunc(t *testing.T) {
	var s Setting = SetFunc(func(v interface{}) error {
		return nil
	})
	s.Set(1)
	s.Set(true)
	s.Set(nil)
}
