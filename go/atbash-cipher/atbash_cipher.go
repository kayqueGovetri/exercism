package atbashcipher

import "strings"

var atbash = map[rune]rune{
	'a': 'z',
	'b': 'y',
	'c': 'x',
	'd': 'w',
	'e': 'v',
	'f': 'u',
	'g': 't',
	'h': 's',
	'i': 'r',
	'j': 'q',
	'k': 'p',
	'l': 'o',
	'm': 'n',
	'n': 'm',
	'o': 'l',
	'p': 'k',
	'q': 'j',
	'r': 'i',
	's': 'h',
	't': 'g',
	'u': 'f',
	'v': 'e',
	'w': 'd',
	'x': 'c',
	'y': 'b',
	'z': 'a',
}

func Atbash(s string) string {
	runes := []rune(strings.ToLower(s))
	response := []rune{}
	count := 0
	for _, letter := range runes {
		value, ok := atbash[letter]
		if !ok {
			if letter >= '0' && letter <= '9' {
				response = append(response, letter)
				count++
			}
			continue
		}

		if count == 5 {
			response = append(response, ' ')
			count = 0
		}
		response = append(response, value)
		count++
	}
	return string(response)
}
