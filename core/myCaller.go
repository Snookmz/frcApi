package core

import (
	"runtime"
	"strings"
)

// MyCaller returns the caller of the function that called it
func MyCaller() string {

	// we get the callers as uintptrs - but we just need 1
	fpcs := make([]uintptr, 1)

	// skip 3 levels to get to the caller of whoever called Caller()
	n := runtime.Callers(3, fpcs)
	if n == 0 {
		return "n/a" // proper error her would be better
	}

	// get the info of the actual function that's in the pointer
	fun := runtime.FuncForPC(fpcs[0] - 1)
	if fun == nil {
		return "n/a"
	}

	fullName := strings.Split(fun.Name(), "/")

	function := fullName[len(fullName) - 1]

	//function = strings.Replace(function, "loggerAbs.go:26", "", 1)

	// return its name
	return function
}