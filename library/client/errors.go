package client

type RpcConnectionError struct {
	Message string `json:"message"`
}

func NewRpcConnectionError(message string) error {
	return &RpcConnectionError{Message: message}
}
func (e *RpcConnectionError) Error() string {
	return e.Message
}

type AdressInvalidError struct {
	Message string `json:"message"`
}

func NewRpcAdressInvalidError(message string) error {
	return &AdressInvalidError{Message: message}
}
func (e *AdressInvalidError) Error() string {
	return e.Message
}

// error rype for Block too far in future
type BlockTooFarInFutureError struct {
	Message string `json:"message"`
}

func NewBlockTooFarInFutureError(message string) error {
	return &BlockTooFarInFutureError{Message: message}
}
func (e *BlockTooFarInFutureError) Error() string {
	return e.Message
}

// error type for Era too far in future
type EraTooFarInFutureError struct {
	Message string `json:"message"`
}

func NewEraTooFarInFutureError(message string) error {
	return &EraTooFarInFutureError{Message: message}
}
func (e *EraTooFarInFutureError) Error() string {
	return e.Message
}

// new invalid era number error
type InvalidEraNumberError struct {
	Message string `json:"message"`
}

func NewInvalidEraNumberError(message string) error {
	return &InvalidEraNumberError{Message: message}
}
func (e *InvalidEraNumberError) Error() string {
	return e.Message
}
