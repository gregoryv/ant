package ant_test

import (
	"fmt"

	"github.com/gregoryv/ant"
)

func ExampleSetting() {
	app := NewApp(Binding(":9090"))
	fmt.Print(app.bind)
	// output:
	// :9090
}

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

type Binding string

func (me Binding) Set(v interface{}) error {
	switch v := v.(type) {
	case *App:
		v.bind = string(me)
	default:
		return ant.SetFailed(v, me)
	}
	return nil
}
