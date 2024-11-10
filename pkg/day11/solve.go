package day11

var SEQUENCE = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h' /*'i',*/, 'j', 'k' /*'l',*/, 'm', 'n' /*'o',*/, 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
var INDEXES map[rune]int = make(map[rune]int)

func init() {
	for index, value := range SEQUENCE {
		INDEXES[value] = index
	}
}

func NextRune(value []rune) []rune {
	incrementing := len(value) - 1
	for true {
		// if we're incrementing position -1, then we need to append
		// a new chacter
		if incrementing == -1 {
			value = append([]rune{'a'}, value...)
		} else if value[incrementing] == 'z' {
			value[incrementing] = 'a'
			incrementing -= 1
			continue
		} else {
			value[incrementing] = SEQUENCE[INDEXES[value[incrementing]]+1]
		}

		break
	}

	return value
}

func ValidPassword(value []rune) bool {
	var firstPair rune
	pairs := false
	sequenced := false

	for index := range value {
		// forbidden character
		_, ok := INDEXES[value[index]]
		if !ok {
			return false
		}

		// test for numbers in sequence
		if !sequenced {
			if index+2 < len(value) {
				sequenced = (value[index] == (value[index+1]-1) && value[index+1] == (value[index+2]-1))
			}
		}

		// check for unique pair, need at least two
		if !pairs {
			if index != len(value)-1 {
				if value[index] == value[index+1] {
					if firstPair == 0 {
						firstPair = value[index]
					} else if firstPair != value[index] {
						pairs = true
					}
				}
			}
		}
	}

	return pairs && sequenced
}

func NextPassword(value string) string {
	current := []rune(value)
	current = NextRune(current)

	for !ValidPassword(current) {
		current = NextRune(current)
	}

	return string(current)
}
