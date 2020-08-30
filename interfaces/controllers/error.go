package controllers

// Error エラー
type Error struct {
	Message string
}

// NewError エラー生成
func NewError(err error) *Error {
	return &Error{
		Message: err.Error(),
	}
}
