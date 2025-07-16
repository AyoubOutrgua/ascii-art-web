package functions

func CheckBanner(banner string) bool {
	return banner == "shadow" || banner == "standard" || banner == "thinkertoy"
}
