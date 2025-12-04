package formatters

import (
	"fmt"
	"runtime"

	"github.com/sirupsen/logrus"
)

func NewJSONFormatter() *logrus.JSONFormatter {
	return &logrus.JSONFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			function, line := customCallerFormatter(f)
			combined := fmt.Sprintf("%s:%d", function, line)

			return combined, ""
		},
	}
}
