package data

import (
	"regexp"
	"unicode"
)

var reArrayPath = regexp.MustCompile(`^([^\[]+)?(\[[0-9]+\])?$`)

// split splits a string describing a JSON path into its components,
// i.e., strings representing one level of descend in a map. Array
// indexes in brackets are returned together with their "parent string".
//
// Examples:
//  split(`["store"]["book"][0]["title"]`)
//  // []string{"store", "book[0]", "title"}
//  split(`store.book[0].title`)
//  // []string{"store", "book[0]", "title"}
func split(s string) []string {
	i := 0
	result := []string{}
	part := ""
	runes := []rune(s)
	length := len(runes)
	for i < length {
		r := runes[i]
		switch r {
		case '\\':
			// if we see a backslash, the next character will be
			// appended to the string verbatim (i.e. we can escape
			// dots, brackets etc outside of brackets). if it is
			// the last character of the input string it will be ignored.
			i++
			if i < length {
				part += string(runes[i])
			}
		case '.':
			// if we see a dot, we will finish the current component
			// and append it to the result list. if the current
			// component is empty, the dot will be ignored.
			if part != "" {
				result = append(result, part)
				part = ""
			}
		case '[':
			// if we see an opening bracket, we can either have
			// hoge[123] or hoge["key"] or some invalid situation.
			if i < length-1 {
				nr := runes[i+1]
				// if we have the hoge["key"] situation, get the
				// part until the closing bracket.
				if nr == '"' || nr == '\'' {
					inbracket := splitBracket(runes, i+2, nr)
					if inbracket != "" {
						if part != "" {
							result = append(result, part)
						}
						result = append(result, inbracket)
						part = ""
						i += 1 + len(inbracket) + 2 // " + inner bracket + "]
						break
					}
				}
			}
			// NB. if we have a string after an opening bracket that
			// does not begin with a quote character (this can be numeric
			// or not), it will be copied verbatim
			part += string(r)
		default:
			part += string(r)
		}
		i++
	}
	if part != "" {
		result = append(result, part)
	}
	return result
}

// splitBracket returns a string in `runes` that begins at position `i`
// and is followed by the given quote character and a closing bracket.
// It can be used to extract `key` from `hoge["key"]`. If there is an
// integer index in brackets following, this will be returned as well.
// If there is no string matching the conditions, returns an empty string.
//
// Example:
//  splitBracket([]rune(`a["hoge"].b`), 3, '"')
//  // `hoge`
//  splitBracket([]rune(`a["hoge"][123]`), 3, '"')
//  // `hoge[123]`
func splitBracket(runes []rune, i int, quote rune) string {
	length := len(runes)
	result := ""
	for i < length {
		r := runes[i]
		// if the current character is the required quote character ...
		if r == quote {
			if i < length-1 {
				// ... and the next character is the closing bracket:
				if runes[i+1] == ']' {
					index := ""
					// if there is a following opening bracket, which
					// may or may not be followed by an array index,
					// try to get that index and append it to what
					// we found so far (it will be an empty string if
					// there is no array index but a string found)
					if i < length-4 && runes[i+2] == '[' {
						index = getArrayIndex(runes, i+3)
					}
					// return what we found until now
					return result + index
				}
			}
		}
		// otherwise just append the found character to the intermediate
		// result
		result += string(r)

		i++
	}
	// if we never fulfill the above condition, return an empty string
	return ""
}

// getArrayIndex returns a string in `runes` that begins at position `i`,
// consists only of digits and is followed by a closing bracket.
// The string that was found is wrapped in brackets before returning.
// It can be used to extract `[123]` from `hoge[123]`. If there is no
// string matching the conditions, returns an empty string.
//
// Example:
//  getArrayIndex([]rune(`hoge[123]`), 5)
//  // `[123]`
func getArrayIndex(runes []rune, i int) string {
	length := len(runes)
	result := ""
	for i < length {
		r := runes[i]
		if r == ']' {
			break
		} else if unicode.IsNumber(r) {
			result += string(r)
		} else {
			return ""
		}
		i++
	}
	if result != "" {
		return "[" + result + "]"
	}
	return ""
}

// scanMap does basically what is described in the Map.Get documentation.
// The value found at p is written to v.
func scanMap(m Map, p string, v *Value) (err error) {
	path, err := CompilePath(p)
	if err != nil {
		return err
	}
	val, err := path.Evaluate(m)
	if err != nil {
		return err
	}
	*v = val
	return nil
}

// setInMap does basically what is described in the Map.Set documentation.
// The value at v is written to m at the given path.
func setInMap(m Map, p string, v Value) (err error) {
	path, err := CompilePath(p)
	if err != nil {
		return err
	}
	return path.Set(m, v)
}
