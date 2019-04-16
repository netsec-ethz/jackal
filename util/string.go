/*
 * Copyright (c) 2018 Miguel Ángel Ortuño.
 * See the LICENSE file for more information.
 */

package util

import "strings"

// SplitKeyAndValue splits a string between 'key' and 'value' sub elements.
func SplitKeyAndValue(str string, sep byte) (key string, value string) {
	j := -1
	for i := 0; i < len(str); i++ {
		if str[i] == sep {
			j = i
			break
		}
	}
	if j == -1 {
		return "", ""
	}
	key = str[0:j]
	value = str[j+1:]
	return
}

// StringRepeat returns n times the specified string separated by the separator value.
func StringRepeat(s string, separator string, n int) string {
	var b strings.Builder
	for i := 0; i < n; i++ {
		if i > 0 {
			b.WriteString(separator)
		}
		b.WriteString(s)
	}
	return b.String()
}
