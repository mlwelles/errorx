package errorx

import (
	"fmt"

	"github.com/pkg/errors"
	"go.uber.org/multierr"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type StackTracer interface {
	StackTrace() errors.StackTrace
}

func Unwrap(err error) error {
	return errors.Unwrap(err)
}

func And(fn1 func(error) bool, fn2 func(error) bool, fns ...func(error) bool) func(error) bool {
	return func(e error) bool {
		if !fn1(e) {
			return false
		}
		if !fn2(e) {
			return false
		}
		for _, fn := range fns {
			if !fn(e) {
				return false
			}
		}
		return true
	}
}

func Or(fn1 func(error) bool, fn2 func(error) bool, fns ...func(error) bool) func(error) bool {
	return func(e error) bool {
		if fn1(e) {
			return true
		}
		if fn2(e) {
			return true
		}
		for _, fn := range fns {
			if fn(e) {
				return true
			}
		}
		return false
	}
}

func WhereNotIs(target error) func(error) bool {
	return func(err error) bool {
		return NotIs(err, target)
	}
}

func WhereIs(target error) func(error) bool {
	return func(err error) bool {
		return Match(err, target)
	}
}

func WhereAs(target error) func(error) bool {
	return func(err error) bool {
		return As(err, target)
	}
}

func WhereNotAs(target error) func(error) bool {
	return func(err error) bool {
		return NotAs(err, target)
	}
}

func WhereAll(fns ...func(error) bool) func(error) bool {
	return WhereErrorAll(fns...)
}

func NotNilNotIs(err, target error) bool {
	return NotNil(err) && NotIs(err, target)
}

func NotNilAndNotAs(err, target error) bool {
	return NotNil(err) && NotAs(err, target)
}

func NotAs(err, target error) bool {
	return !As(err, target)
}

func NotIs(err error, targets ...error) bool {
	return !Match(err, targets...)
}

func IsOrAs(err error, targets ...error) bool {
	for _, cmp := range targets {
		switch {
		case err == nil && cmp == nil:
			return true
		case errors.Is(err, cmp):
			return true
		case SafelyAs(err, cmp):
			return true
		}
	}
	return false
}

func Code(err error) codes.Code {
	c := status.Code(err)
	if c == codes.Unknown {
		st, _ := status.FromError(err)
		if st != nil {
			c = st.Code()
		}
	}
	if c == codes.Unknown {
		causalError := errors.Cause(err)
		st, _ := status.FromError(causalError)
		if st != nil {
			c = st.Code()
		}
	}
	return c
}

// Is reports whether any error in err's chain matches target.
//
// The chain consists of err itself followed by the sequence of errors obtained by
// repeatedly calling Unwrap.
//
// An error is considered to match a target if it is equal to that target or if
// it implements a method Is(error) bool such that Is(target) returns true.
func Is(err error, target error) bool {
	return errors.Is(err, target)
}

// Match improved "Is" -- combines functionality of "Is" and (Safely) "As" but also recursively unwraps the error itself, understands
// wrapped grpc errors as well, is nil safe, allows multiple comparisons.
func Match(e error, targets ...error) bool {
	if len(targets) == 0 {
		return e != nil
	}
	for _, err := range multierr.Errors(e) {
		for _, tErr := range targets {
			for _, t := range multierr.Errors(tErr) {
				if err == nil {
					if t == nil {
						return true
					}
				} else if t != nil {
					if match(err, t, false) {
						return true
					}
				}
			}
		}
	}
	return false
}

func Flatten(errs ...error) ErrorSlice {
	flattened := NewErrorSlice()
	for _, e := range errs {
		flattened = append(flattened, Errors(e)...)
	}
	return flattened
}

func Split(err error) ErrorSlice {
	return Errors(err)
}

func Errors(err error) ErrorSlice {
	return split(err, NewErrorSlice())
}

func split(err error, errs ErrorSlice) ErrorSlice {
	if err == nil {
		return errs
	}
	errs = append(errs, err)
	partErrs := multierr.Errors(err)
	e1 := Unwrap(err)
	if e1 != nil && e1 != err {
		partErrs = append(partErrs, e1)
	}
	e2 := Cause(err)
	if e2 != nil && e2 != err {
		partErrs = append(partErrs, e1)
	}
	for _, e := range partErrs {
		if e == err {
			continue
		}
		errs = split(e, errs)
	}
	return errs
}

//Intersect improved "Match" -- combines functionality of "Is" and (Safely) "As" but also digs into the error itself, understands
// wrapped grpc errors as well, is nil safe, allows multiple comparisons. will unwrap the err and targets recursively and
// return true if there's any intersections
func Intersect(err error, targets ...error) bool {
	expandedTargets := NewErrorSlice(targets...)
	for _, t := range targets {
		expandedTargets = append(expandedTargets, Errors(t)...)
	}
	targets = expandedTargets
	if len(targets) == 0 {
		return err != nil
	}
	for _, t := range targets {
		if err == nil {
			if t == nil {
				return true
			}
		} else if t != nil {
			if match(err, t, true) {
				return true
			}
		}
	}
	return false
}

func IsGRPC(err error) bool {
	_, ok := status.FromError(err)
	return ok
}

func match(err error, target error, recurseTarget bool) bool {
	switch {
	case err == nil, target == nil:
		return err == target
	case err.Error() == target.Error():
		return true
	case Is(err, target), SafelyAs(err, target):
		return true
	case IsGRPC(err):
		return matchGRPC(err, target)
	default:
		for _, cause := range Causes(err) {
			if match(cause, target, false) {
				return true
			}
			if recurseTarget {
				for _, t := range Causes(target) {
					if match(cause, t, false) {
						return true
					}
				}
			}
		}
		return false
	}
}

func matchGRPC(err, target error) bool {
	// see if its a grpq status error and we can just
	// match the error from the message
	st, ok := status.FromError(err)
	if ok {
		msg := st.Message()
		if msg == target.Error() {
			return true
		} else if Match(fmt.Errorf(st.Message()), target) {
			return true
		}
	}
	// see if it's grpc status error and try and unwrap from the code
	causalError := err
	c := status.Code(err)
	if c == codes.Unknown {
		causalError = errors.Cause(err)
		st, _ := status.FromError(causalError)
		if st != nil {
			c = st.Code()
		}
	}
	if c != codes.Unknown {
		if causalError != nil {
			if errors.Is(err, causalError) {
				return true
			}
			causalError = errors.Cause(causalError)
			if causalError != nil {
				if errors.Is(err, causalError) {
					return true
				}
			}
		}

	}
	return false
}

type MultiError interface {
	error
	Errors() []error
}

// Cause returns the underlying cause of the error, if possible.
// An error value has a cause if it implements the following
// interface:
//
//     type causer interface {
//            Cause() error
//     }
//
// If the error does not implement Cause, the original error will
// be returned. If the error is nil, nil will be returned without further
// investigation.

type Causer interface {
	error
	Cause() error
}

func Cause(err error) error {
	for err != nil {
		cause, ok := err.(Causer)
		if !ok {
			break
		}
		err = cause.Cause()
	}
	return err
}

func Causes(err error) []error {
	var errs []error
	multi, ok := err.(MultiError)
	if ok {
		for _, e := range multi.Errors() {
			errs = append(errs, Causes(e)...)
		}
	}
	for err != nil {
		causer, ok := err.(Causer)
		if !ok {
			break
		}
		cause := causer.Cause()
		if cause != nil && cause != err && cause.Error() != err.Error() {
			errs = append(errs, cause)
		}
		err = cause
	}
	return errs
}

// The chain consists of err itself followed by the sequence of errors obtained by
// repeatedly calling Unwrap.
//
// An error matches target if the error's concrete value is assignable to the value
// pointed to by target, or if the error has a method As(interface{}) bool such that
// As(target) returns true. In the latter case, the As method is responsible for
// setting target.
//
// As will panic if target is not a non-nil pointer to either a type that implements
// error, or to any interface type. As returns false if err is nil.
func As(err error, target interface{}) bool {
	if err == nil && target != nil {
		return false
	}
	return errors.As(err, target)
}

// Same as as but will return false instead of panic if target is not a non-nil pointer to either a type that implements error, or to any interface type.
func SafelyAs(err error, target interface{}) (okAs bool) {
	defer func() {
		if err := recover(); err != nil {
			okAs = false
		}
	}()
	if err == nil && target != nil {
		okAs = false
		return okAs
	}
	okAs = As(err, target)
	return okAs

}
func Nil(err error) bool {
	return err == nil
}

func NotNil(err error) bool {
	return err != nil
}

// Errorf formats according to a format specifier and returns the string
// as a value that satisfies error.
// Errorf also records the stack trace at the point it was called.
func Errorf(format string, args ...interface{}) error {
	return errors.Errorf(format, args...)
}

// New returns an error with the supplied message.
// New also records the stack trace at the point it was called.
func New(str string, errs ...error) error {
	errs = ErrorSlice(errs).Filter(NotNil)
	switch len(errs) {
	case 0:
		return errors.New(str)
	case 1:
		return Wrap(errs[0], str)
	default:
		errs[0] = Wrap(errs[0], str)
		return Combine(errs...)
	}
}

// NewCombine creates new error combined with any passed as args
// but without wrapping stack trace
func NewCombined(msg string, errs ...error) error {
	return Combine(append([]error{fmt.Errorf(msg)}, errs...)...)
}

// use to create constatns without stack traces
func new(msg string, errs ...error) error {
	return NewCombined(msg, errs...)
}

func Combinef(err error, msg string, args ...interface{}) error {
	return Combine(err, fmt.Errorf(msg, args...))
}

// Wrap returns an error annotating err with a stack trace
// at the point Wrap is called, and the supplied message.
// If err is nil, Wrap returns nil.
func Wrap(err error, message string, extra ...error) error {
	err = errors.Wrap(err, message)
	if len(extra) == 0 {
		return err
	}
	return Combine(append([]error{err}, extra...)...)
}

func WrapError(wrapped, wrapper error) error {
	return Combine(wrapped, wrapper)
}

// Wrapf returns an error annotating err with a stack trace
// at the point Wrapf is called, and the format specifier.
// If err is nil, Wrapf returns nil.
func Wrapf(err error, format string, args ...interface{}) error {
	return errors.Wrapf(err, format, args...)
}

// Combine combines the passed errors into a single error.
//
// If zero arguments were passed or if all items are nil, a nil error is
// returned.
//
// 	Combine(nil, nil)  // == nil
//
// If only a single error was passed, it is returned as-is.
//
// 	Combine(err)  // == err
//
// Combine skips over nil arguments so this function may be used to combine
// together errors from operations that fail independently of each other.
//
// 	multierr.Combine(
// 		reader.Close(),
// 		writer.Close(),
// 		pipe.Close(),
// 	)
//
// If any of the passed errors is a multierr error, it will be flattened along
// with the other errors.
//
// 	multierr.Combine(multierr.Combine(err1, err2), err3)
// 	// is the same as
// 	multierr.Combine(err1, err2, err3)
//
// The returned error formats into a readable multi-line error message if
// formatted with %+v.
//
// 	fmt.Sprintf("%+v", multierr.Combine(err1, err2))
func Combine(errors ...error) error {
	return multierr.Combine(errors...)
}

// Append appends the given errors together. Either value may be nil.
//
// This function is a specialization of Combine for the common case where
// there are only two errors.
//
// 	err = multierr.Append(reader.Close(), writer.Close())
//
// The following pattern may also be used to record failure of deferred
// operations without losing information about the original error.
//
// 	func doSomething(..) (err error) {
// 		f := acquireResource()
// 		defer func() {
// 			err = multierr.Append(err, f.Close())
// 		}()
func Append(left error, right error) error {
	return multierr.Append(left, right)
}

func CauseOrError(err error) string {
	if err == nil {
		return ""
	}
	cause := ""
	ec := Cause(err)
	if ec != nil {
		cause = ec.Error()
	}
	return cause
}

type StackTrace = errors.StackTrace

func GetStackTrace(err error) errors.StackTrace {
	stackErr, ok := err.(StackTracer)
	if !ok {
		return nil
	}
	return stackErr.StackTrace()
}

func StackTraceString(err error) string {
	trace := GetStackTrace(err)
	if trace == nil {
		return ""
	}
	return fmt.Sprintf("%+v", trace)
}
