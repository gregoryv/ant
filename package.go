// Package ant provides a setting interface.
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
// type of value.
func SetFailed(v interface{}, setting Setting) error {
	return fmt.Errorf("failed to set %T on %T", setting, v)
}
