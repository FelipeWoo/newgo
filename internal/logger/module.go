package logger

var currentModule string

func SetModule(name string) {
	currentModule = name
}

func getModule() string {
	if currentModule == "" {
		return "unknown"
	}
	return currentModule
}
