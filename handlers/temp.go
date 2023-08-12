package handlers

var test map[string]any = make(map[string]any)

func Get(key string) any {
	return test[key]
}

func Set(key string, value any) {
	test[key] = value
}
