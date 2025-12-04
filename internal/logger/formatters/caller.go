package formatters

import (
	"runtime"
	"runtime/debug"
	"strings"
	"sync"
)

var (
	buildPath     string
	buildPathOnce sync.Once
)

func init() {
	buildPathOnce.Do(func() {
		if info, ok := debug.ReadBuildInfo(); ok && info.Main.Path != "" {
			buildPath = info.Main.Path + "/"
		}
	})
}

// customCallerFormatter returns the function name and line number
//
// this information is only available if the binary was built with go modules
func customCallerFormatter(f *runtime.Frame) (string, int) {
	function := f.Function
	if buildPath != "" {
		function = strings.ReplaceAll(f.Function, buildPath, "")
	}

	function = function + "()"

	return function, f.Line
}
