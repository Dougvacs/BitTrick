package compression

import (
	"strconv"
)

func Rle(s string) string {
	var cs = ""
	var duplicate_counter = 0
	var last_char = string([]byte(s)[0])
	for i := range s {
		if string(s[i]) == last_char {
			duplicate_counter += 1
		} else {
			cs += strconv.Itoa(duplicate_counter) + last_char
			duplicate_counter = 1
			last_char = string(s[i])
		}
	}
	cs += strconv.Itoa(duplicate_counter) + last_char
	return cs
}

func RleDecode(cs string) string {
	var s = ""
	var duplicate_counter = 0
	for i := range cs {
		parsed_int, conv_error := strconv.Atoi(string(cs[i]))
		if conv_error != nil {
			for j := 0; j < duplicate_counter; j++ {
				s += string(cs[i])
			}
		} else {
			duplicate_counter = parsed_int
		}
	}
	return s
}
