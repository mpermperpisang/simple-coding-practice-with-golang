package functions

import "strings"

func JoinArray(array []string) string {
	return strings.Join(array, ", ") + "]"
}
