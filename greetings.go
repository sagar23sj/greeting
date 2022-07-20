package greeting

import (
	"fmt"

	"github.com/fatih/color"
)

func Hello(name string) string {
	color.Cyan("welcome %s", name)
	msg := fmt.Sprintf("Hi, %v. Welcome!", name)
	return msg
}
