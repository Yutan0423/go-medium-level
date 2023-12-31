package apperrors

type ErrCode string

const (
	UnKnown ErrCode = "U000"

	InsertDataFailed ErrCode = "S001"
	GetDataFaild     ErrCode = "S002"
	NAData           ErrCode = "S003"
	NoTargetData     ErrCode = "S004"
	UpdateDataFailed ErrCode = "S005"

	ReqBodyDecodeFailed ErrCode = "R001"
	BadParam            ErrCode = "R002"

	RequiredAuthorizationHeader ErrCode = "A001"
	CannotMakeValidator         ErrCode = "A002"
	Unauthorized                ErrCode = "A003"
)

func (code ErrCode) Wrap(err error, message string) *MyAppError {
	return &MyAppError{ErrCode: code, Message: message, Err: err}
}
