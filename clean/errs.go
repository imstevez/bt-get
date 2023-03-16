package clean

import "strings"

type Errs []error

func (err Errs) Error() string {
	strs := make([]string, len(err))
	for i, e := range err {
		strs[i] = e.Error()
	}
	return strings.Join(
		strs, `; `,
	)
}
