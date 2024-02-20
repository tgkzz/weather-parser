package model

type constError string

func (err constError) Error() string {
	return string(err)
}

const (
	ErrNoInternetConnection          = constError("service dont have internet connection")
	ErrEmptyParams                   = constError("some of the parameters are missing")
	ErrNoCity                        = constError("no city in the database")
	ErrToManyRequestToAnotherService = constError("too many request was sent to another service")
)
