package util

func IsRedirectType(args string) (string, bool) {
	if args == "<" {
		return "stdin", false
	} else if args == "<<" {
		return "stdin", true
	} else if args == ">" {
		return "stdout", false
	} else if args == ">>" {
		return "stdout", true
	}
	return "", false
}
