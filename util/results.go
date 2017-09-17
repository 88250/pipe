package util

// Result.
type Result struct {
	Succ bool        `json:"succ"` // successful or not
	Code string      `json:"code"` // return code
	Msg  string      `json:"msg"`  // message
	Data interface{} `json:"data"` // data object
}

// NewResult creates a result with Succ=true, Code="0", Msg="", Data=nil.
func NewResult() *Result {
	return &Result{
		Succ: true,
		Code: "0",
		Msg:  "",
		Data: nil,
	}
}
