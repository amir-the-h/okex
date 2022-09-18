package rest

import (
	"encoding/json"
	"io"
)

type res interface {
	ErrStatus() error
}

func Decode(Body io.ReadCloser, v res) error {
	d := json.NewDecoder(Body)
	err := d.Decode(&v)
	if err != nil {
		return err
	}
	return v.ErrStatus()
}
