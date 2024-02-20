package model

type constError string

func (err constError) Error() string {
	return string(err)
}

const (
	Err4xxStatusCode = constError("another service is receiving ...")
	ErrEmptyParams   = constError("some of the parameters are missing")
)
