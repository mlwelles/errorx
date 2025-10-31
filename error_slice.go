package errorx

import "fmt"

func (slice ErrorSlice) Errors() []string {
	return slice.MapString(func(e error) string { return e.Error() })
}
func (slice ErrorSlice) Quoted() []string {
	return slice.MapString(func(e error) string { return fmt.Sprintf("%q", e.Error()) })
}

func (slice ErrorSlice) Combined() error {
	return Combine(slice...)
}

func (slice ErrorSlice) CombinedDistinct() error {
	var err error
	for _, e := range slice {
		if err == nil || !Match(err, e) {
			err = Combine(err, e)
		}
	}
	return err
}
