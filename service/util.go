package service

func WrapJson(name string, item interface{}) map[string]interface{} {
	return map[string]interface{}{
		name: item,
	}
}
