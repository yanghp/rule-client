package client

import (
	"errors"
)

type MuteMap struct {
	mp map[string]interface{}
}

func Provider(mp map[string]interface{}) *MuteMap {
	return &MuteMap{mp: mp}
}

// ReadBytes is not supported by the env provider.
func (e *MuteMap) ReadBytes() ([]byte, error) {
	return nil, errors.New("mutemap provider does not support this method")
}

// Read returns the loaded map[string]interface{}.
func (e *MuteMap) Read() (map[string]interface{}, error) {
	return e.mp, nil
}

// Watch is not supported.
func (e *MuteMap) Watch(cb func(event interface{}, err error)) error {
	return errors.New("mutemap provider does not support this method")
}
