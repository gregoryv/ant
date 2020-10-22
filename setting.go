// Package ant provides a setting interface.
package ant

// Setting is anything that can be set on some entity
type Setting interface {
	Set(v interface{}) error
}

type SetFunc func(v interface{}) error

func (me SetFunc) Set(v interface{}) error {
	return me(v)
}
