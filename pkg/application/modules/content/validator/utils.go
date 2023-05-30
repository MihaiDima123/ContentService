package validator

func valueLessThan(value string, toCompare int) bool {
	return len(value) <= toCompare
}
