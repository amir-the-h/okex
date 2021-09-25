package responses

type (
	Basic struct {
		Code int    `json:"code,string"`
		Msg  string `json:"msg,omitempty"`
	}
)
