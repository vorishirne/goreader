package structures

func SelectKey(arr *[]map[string]interface{}, key, value string) *map[string]interface{} {
	for _, kv := range *arr {
		if v, ok := kv[key].(string); ok && v == value {
			return &kv
		}
	}
	return nil
}

func FilterKey(arr *[]map[string]interface{}, key, value string) *[]map[string]interface{} {
	filteredArr := make([]map[string]interface{}, 0)
	for _, kv := range *arr {
		if v, ok := kv[key].(string); ok && v == value {
			filteredArr = append(filteredArr, kv)
		}
	}
	return &filteredArr
}
