// Package ant provides a setting interface.
package ant

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
