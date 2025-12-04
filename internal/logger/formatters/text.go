package formatters

import (
	"fmt"
	"runtime"

	"github.com/sirupsen/logrus"
)

func NewTextFormatter() *logrus.TextFormatter {
	return &logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			function, line := customCallerFormatter(f)
			combined := fmt.Sprintf("%s:%d", function, line)

			return combined, ""
		},
	}
}
