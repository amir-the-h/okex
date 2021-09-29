package events

import (
	"encoding/json"
	"github.com/amir-the-h/okex"
)

type (
	Basic struct {
		ID    string         `json:"id,omitempty"`
		Event string         `json:"event"`
		Code  int            `json:"code,omitempty,string"`
		Msg   string         `json:"msg,omitempty"`
		Op    okex.Operation `json:"op,omitempty"`
		Arg   *Argument      `json:"arg,omitempty"`
		Args  []*Argument    `json:"args,omitempty"`
		Data  []*Argument    `json:"data,omitempty"`
	}
	Argument struct {
		arg        map[string]interface{}
		untypedArg []interface{}
	}
	Success struct {
		Code int            `json:"code,omitempty,string"`
		Msg  string         `json:"msg,omitempty"`
		ID   string         `json:"id,omitempty"`
		Op   okex.Operation `json:"op,omitempty"`
		Data []*Argument    `json:"data,omitempty"`
	}
	Error struct {
		Event string         `json:"event,omitempty"`
		Msg   string         `json:"msg,omitempty"`
		Op    string         `json:"op,omitempty"`
		Code  okex.JSONInt64 `json:"code"`
		Args  []*Argument    `json:"args,omitempty"`
		Arg   *Argument      `json:"arg,omitempty"`
		Data  []*Argument    `json:"data,omitempty"`
		ID    string         `json:"id,omitempty"`
	}
	Login struct {
		Event string `json:"event"`
		Code  string `json:"code"`
		Msg   string `json:"msg"`
	}
	Subscribe struct {
		Event string    `json:"event"`
		Arg   *Argument `json:"arg"`
	}
	Unsubscribe struct {
		Event string    `json:"event"`
		Arg   *Argument `json:"arg"`
	}
)

func (a *Argument) Get(k string) (interface{}, bool) {
	v, ok := a.arg[k]
	return v, ok
}

func (a *Argument) UnmarshalJSON(buf []byte) error {
	a.arg = make(map[string]interface{})
	if json.Unmarshal(buf, &a.arg) != nil {
		return json.Unmarshal(buf, &a.untypedArg)
	}

	return nil
}
