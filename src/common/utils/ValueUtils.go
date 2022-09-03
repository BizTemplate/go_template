package utils

type ValueUtils struct {
}

func (u *ValueUtils) GetOrDefault(value string, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return value
}
