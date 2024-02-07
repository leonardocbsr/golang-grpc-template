package formatters

import (
	"runtime"
	"runtime/debug"
	"strings"
)

// customCallerFormatter returns the function name and line number
//
// this information is only available if the binary was built with go modules
func customCallerFormatter(f *runtime.Frame) (string, int) {
	info, _ := debug.ReadBuildInfo()
	path := info.Main.Path + "/"

	var function string
	if info.Main.Path != "" {
		function = strings.ReplaceAll(f.Function, path, "")
	}

	function = function + "()"

	return function, f.Line
}
