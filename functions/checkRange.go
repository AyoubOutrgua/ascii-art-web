package functions

func CheckRange(inputText string) bool {
	for _, v := range inputText{
		if (v < 32 || v > 126) && v!='\n' && v!='\r'{
			return false
		}
	}
	return true
}