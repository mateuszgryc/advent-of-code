package common

func ConvertRuneToPriority(r rune) int {
	num := int(r%97) + 1
	if num > 27 {
		num -= 39
	}
	return num
}
