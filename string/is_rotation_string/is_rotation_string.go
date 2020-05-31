package chapter01

import "strings"

// IsRotationString check if string a is a rotation of another string b
func IsRotationString(a, b string) bool {
	return len(a) > 0 && len(a) == len(b) && strings.Index(a+a, b) != -1
}
