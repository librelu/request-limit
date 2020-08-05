package utilerrors

import (
	"errors"
	"fmt"

	"github.com/ztrue/tracerr"
)

func Wrap(err error, msg string) error {
	e := fmt.Errorf("%s:%w", msg, err)
	return tracerr.Wrap(e)
}

func New(msg string) error {
	return tracerr.Wrap(errors.New(msg))
}
