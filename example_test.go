package ant

import (
	"fmt"
)

func ExampleSetting() {
	app := NewApp(Binding(":9090"))
	fmt.Print(app.bind)
	// output:
	// :9090
}

func NewApp(settings ...Setting) *App {
	app := &App{
		bind: ":8080",
	}
	if err := app.Use(settings...); err != nil {
		panic(err)
	}
	return app
}

type App struct {
	bind string
}

// Use applies the given settings on this service.
func (me *App) Use(settings ...Setting) error {
	for _, setting := range settings {
		if err := setting.Set(me); err != nil {
			return err
		}
	}
	return nil
}

type Binding string

func (me Binding) Set(v interface{}) error {
	switch v := v.(type) {
	case *App:
		v.bind = string(me)
	default:
		return setErr(me, v)
	}
	return nil
}

func setErr(setting, v interface{}) error {
	return fmt.Errorf("%t cannot be set on %t", setting, v)
}
