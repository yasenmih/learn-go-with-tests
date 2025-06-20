package iteration

import "strings"

func Repeat(character string, times int) string {
	var repeated strings.Builder

	//also for i := range times {
	for i := 0; i < times; i++ {
		repeated.WriteString(character)
	}

	return repeated.String()
}
