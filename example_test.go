package ant_test

import (
	"fmt"

	"github.com/gregoryv/ant"
)

func ExampleSetting() {
	// No settings, using default
	fmt.Println(NewApp().bind)

	setting := Binding(":9090")
	fmt.Println(NewApp(setting).bind)
	// output:
	// :8080
	// :9090
}

// NewApp returns a demo application with optional settings applied.
// Panics if invalid settings are provided.
func NewApp(settings ...ant.Setting) *App {
	app := &App{
		bind: ":8080", // default
	}
	ant.MustConfigure(app, settings...)
	return app
}

type App struct {
	bind string
}

// Binding is a setting for configuring a bind value.
type Binding string

// Set implements ant.Setting and can only be applied to *App
func (me Binding) Set(v interface{}) error {
	switch v := v.(type) {
	case *App:
		v.bind = string(me)
	default:
		return ant.SetFailed(v, me)
	}
	return nil
}
