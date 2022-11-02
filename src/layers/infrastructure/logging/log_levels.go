package logging

func activatorLevels(levels []string) levelActivate {
	activate := levelActivate{}

	for _, level := range levels {
		switch level {
		case "all":
			activate.INFO = true
			activate.ERROR = true
			activate.FATAL = true
			return activate
		case "info":
			activate.INFO = true
		case "error":
			activate.ERROR = true
		case "fatal":
			activate.FATAL = true
		}
	}
	return activate
}
