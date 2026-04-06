package runtime

import "runtime"

func GetCallerFunction(level int) (string, int, bool) {

	_, callerFile, callerLine, ok := runtime.Caller(level + 1)
	if ok {
		return callerFile, callerLine, true
	}

	return "", 0, false
}
