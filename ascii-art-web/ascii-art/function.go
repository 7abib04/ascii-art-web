package asciiart

import (
	"errors"
	"os"
	"strings"
)

func Ascii(txt, format string) (string, error) {
	str := ""
	//check if there is carridge return and replace it with \n and then split it
	txt = strings.ReplaceAll(txt, "\r\n", "\n")
	textSlice := strings.Split(txt, "\n")
	// use this function to check is all charcters are valid
	if !charValidation(txt) {
		return "", errors.New("error : invalid char")
	}

	//read from the file
	file, err := os.ReadFile("ascii-art/" + format + ".txt")
	if err != nil {
		return "", errors.New(" error : reading file")
	}
	//splite all file lines in slice
	slice := strings.Split(string(file), "\n")

	for j, txt := range textSlice {
		if txt != "" {
			for i := 0; i < 8; i++ {
				for _, v := range txt {
					firstLine := int(v-32)*9 + 1 + i
					str += slice[firstLine]
				}
				str += "\n"
			}
		} else if j != len(textSlice)-1 {
			str += "\n"
		}
	}
	return str, nil
}

// charValidation checks if all characters are within the printable ASCII range
func charValidation(str string) bool {
	slice := []rune(str)
	for _, char := range slice {
		if (char < 32 || char > 126) && char != '\n' {
			return false
		}
	}
	return true
}
