package ant

import (
	"fmt"
	"testing"
)

func TestConfigure_badSetting(t *testing.T) {
	if err := Configure(nil, &badSetting{}); err == nil {
		t.Error("should fail")
	}
}

func TestMustConfigure(t *testing.T) {
	defer func() {
		if e := recover(); e == nil {
			t.Error("should panic")
		}
	}()
	MustConfigure(nil, &badSetting{})
}

func ExampleConfigure_badSettings() {
	err := Configure(1, &badSetting{})
	fmt.Println(err)
	// output:
	// failed to set ant.badSetting on int
}

type badSetting struct{}

func (me badSetting) Set(v interface{}) error { return SetFailed(v, me) }
