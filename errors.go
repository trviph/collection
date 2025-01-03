package collection

type ErrIsEmpty struct {
	msg string
}

func (err *ErrIsEmpty) Error() string {
	if len(err.msg) > 0 {
		return err.msg
	}
	return "is empty"
}

type ErrNotFound struct {
	msg string
}

func (err *ErrNotFound) Error() string {
	if len(err.msg) > 0 {
		return err.msg
	}
	return "not found"
}

type ErrIndexOutOfRange struct {
	msg string
}

func (err *ErrIndexOutOfRange) Error() string {
	if len(err.msg) > 0 {
		return err.msg
	}
	return "index is out of range"
}
