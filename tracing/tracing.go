package tracing

var (
	trace = false
	key = ""
)

func IsTracing() bool {
	return trace
}

func SetTracing(enabled bool) {
	trace = enabled
}

func SetTracingKey(setKey string) {
	key = setKey
}

func GetTracingKey() string {
	return key
}
