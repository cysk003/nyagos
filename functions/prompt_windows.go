package functions

import (
	"github.com/zetamatta/go-windows-su"
)

func isElevated() bool {
	val, _ := su.IsElevated()
	return val
}
