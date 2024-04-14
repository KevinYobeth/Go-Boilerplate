package utils

func ValueOrEmptyString(value *string) string {
	if value == nil {
		return ""
	}

	return *value
}
