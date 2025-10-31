package errorx

import (
	"fmt"
	"github.com/noho-digital/logging"
)

type Log int

const (
	LogNever Log = iota
	LogError
	LogInfo
)

func LoggingWrapper(level Log, logger logging.Logger, while string, args ...interface{}) func(error, ...interface{}) error {
	if len(args) > 0 {
		while = fmt.Sprintf(while, args...)
	}
	if logger == nil {
		level = LogNever
	}
	if logger != nil && level >= LogInfo {
		logger.Info(while)
	}
	var whileErr error
	if while != "" {
		whileErr = New("error " + while)
	}
	return func(err error, args ...interface{}) error {
		wrapErr := whileErr
		msg := ""
		var msgErr error
		if len(args) > 0 {
			str, ok := args[0].(string)
			if ok {
				msg = str
			}
			args = args[1:]
		}
		if msg != "" {
			if len(args) == 0 {
				msgErr = New(msg)
			} else {
				msgErr = Errorf(msg, args...)
			}
			wrapErr = Combine(wrapErr, msgErr)
		}
		if wrapErr != nil {
			err = Wrap(err, wrapErr.Error())

		}
		if logger != nil && level >= LogError {
			logger.Error(err)
		}
		return err
	}
}
