package validatorImpl

func ValueLessThan(value string, toCompare int) bool {
	return len(value) <= toCompare
}
