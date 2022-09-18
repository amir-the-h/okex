package responses

import "fmt"

type (
	Basic struct {
		Code int    `json:"code,string"`
		Msg  string `json:"msg,omitempty"`
	}
)

func (b *Basic) ErrStatus() error {
	if b.Code != 0 {
		return fmt.Errorf("recevied error:%+v", b)
	}
	return nil
}
