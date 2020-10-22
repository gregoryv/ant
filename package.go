/*
Package ant provides a setting interface.

Apply settings with the configure funcs

   var (
       x X
       setting1 ant.Setting
       setting2 ant.Setting
   )
   err := ant.Configure(&x, setting1, setting2)

or

   ant.MustConfigure(&x, setting1, setting2)

*/
package ant

import "fmt"

// Configure any value with the given settings.
// Stops if applying a setting fails, returning the error.
func Configure(v interface{}, settings ...Setting) error {
	for _, setting := range settings {
		if err := setting.Set(v); err != nil {
			return err
		}
	}
	return nil
}

// MustConfigure panics if a setting fails to apply.
func MustConfigure(v interface{}, settings ...Setting) {
	if err := Configure(v, settings...); err != nil {
		panic(err)
	}
}

// SetFailed returns a preformated error stating type of setting and
// type of value. Use it in settings for normalized setting errors.
func SetFailed(v interface{}, setting Setting) error {
	return fmt.Errorf("failed to set %T on %T", setting, v)
}
