package errorcode


// CodeErr 错误码与信息。共5位，10000 为未知的系统错误。
type CodeErr struct {
	Code int
	Msg  string
}

func (ce CodeErr) Error() string {
	return ce.Msg
}

var (
	// 成功。
	Succ = CodeErr{Code: 0, Msg: "ok"}

	// 其他。
	ErrEmployeeIDEmpty    = CodeErr{Code: 9000, Msg: "Employee id is empty."}
	ErrGinShouldBind      = CodeErr{Code: 9001, Msg: "Gin should bind failed"}
	ErrGinShouldBindQuery = CodeErr{Code: 9002, Msg: "Gin should bind query failed"}
	ErrResourceIDIsNil    = CodeErr{Code: 9003, Msg: "Resource id is nil."}
	ErrMkdirAll           = CodeErr{Code: 9004, Msg: "Mkdir all failed!"}
	ErrSaveFileFailed     = CodeErr{Code: 9005, Msg: "Save file failed!"}
	ErrReadDir            = CodeErr{Code: 9006, Msg: "Read directory failed!"}
	ErrCalcFileMD5        = CodeErr{Code: 9007, Msg: "Calculate file MD5 failed!"}
	ErrUnzipFile          = CodeErr{Code: 9008, Msg: "Unzip file failed!"}

	// 通用。
	ErrBadReqParams   = CodeErr{Code: 10000, Msg: "Req params are invalid."}
	ErrInvalidToken   = CodeErr{Code: 10001, Msg: "Token is invalid"}
	ErrTimeout        = CodeErr{Code: 10002, Msg: "Req process timeout"}
	ErrRecordNotFound = CodeErr{Code: 10003, Msg: "Record not found"}
	ErrCreateMySQL    = CodeErr{Code: 10004, Msg: "Create record to mysql failed"}
	ErrReadMySQL      = CodeErr{Code: 10005, Msg: "Read record from mysql failed."}
	ErrUpdateMySQL    = CodeErr{Code: 10006, Msg: "Update record to mysql failed"}
	ErrDeleteMySQL    = CodeErr{Code: 10007, Msg: "Delete record from mysql failed"}
	ErrJsonMarshal    = CodeErr{Code: 10008, Msg: "JSON marshal failed!"}
	ErrJsonUnmarshal  = CodeErr{Code: 10009, Msg: "JSON unmarshal failed!"}
)