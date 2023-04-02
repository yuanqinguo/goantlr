package goantlr

// Errno 定义错误码
type Errno struct {
	Code    int
	Message string
}

func (err Errno) Error() string {
	return err.Message
}

func (err Errno) WithMsg(msg string) Errno {
	if err.Message == "" {
		err.Message = msg
		return err
	}
	err.Message = err.Message + ", " + msg
	return err
}

func (err Errno) WithErr(rawErr error) Errno {
	var msg string
	if err.Message != "" {
		msg = err.Message + ", "
	}
	if rawErr != nil {
		msg = msg + rawErr.Error()
	}
	err.Message = msg
	return err
}

func (err Errno) NotOK() bool {
	return err.Code != Pass.Code
}

var (
	Pass            = Errno{Code: 0, Message: "pass"}
	SyntaxErr       = Errno{Code: -1, Message: "syntax error"}
	ParamMissErr    = Errno{Code: 1, Message: "parameter missing"}
	CondRequiredErr = Errno{Code: 2, Message: "conditional required"}
	CondEmptyErr    = Errno{Code: 3, Message: "conditional is empty"}
	ParamEmptyErr   = Errno{Code: 4, Message: "parameter is empty"}
	CondMissErr     = Errno{Code: 5, Message: "conditional is missing"}
	ParamErr        = Errno{Code: 6, Message: "parameter error"}
)
