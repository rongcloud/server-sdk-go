package rongcloud

func StringPtr(s string) *string {
	return &s
}

func StringValue(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

func IntPtr(i int) *int {
	return &i
}

func IntValue(i *int) int {
	if i != nil {
		return *i
	}
	return 0
}

func BoolPtr(b bool) *bool {
	return &b
}

func BoolValue(b *bool) bool {
	if b != nil {
		return *b
	}
	return false
}
