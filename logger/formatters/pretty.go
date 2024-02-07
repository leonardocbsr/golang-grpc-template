package formatters

import (
	"fmt"
	"runtime"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

// NewPrettyFormatter returns a pretty formatter
//
// it supports colors and prints the function name and line number
func NewPrettyFormatter() logrus.Formatter {
	return &nested.Formatter{
		HideKeys:    false,
		CallerFirst: false,
		CustomCallerFormatter: func(f *runtime.Frame) string {
			function, line := customCallerFormatter(f)
			combined := fmt.Sprintf(" \x1b[33;1m(%s:%d)\x1b[0m", function, line)

			return combined
		},
	}
}
