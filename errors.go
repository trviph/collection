package collection

import "fmt"

var (
	ErrIsEmpty         error = fmt.Errorf("is empty")
	ErrNotFound        error = fmt.Errorf("not found")
	ErrIndexOutOfRange error = fmt.Errorf("index is out of range")
)
