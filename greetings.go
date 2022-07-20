package greeting

import (
	"fmt"

	"github.com/fatih/color"
)

func Hello(name string) string {
	color.Cyan("welcome %s", name)
	msg := fmt.Sprint("Hi, %s. Welcome!", name)
	return msg
}
