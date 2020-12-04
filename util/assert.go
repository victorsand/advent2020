package util

func AssertTrue(input bool) {
	if !input {
		panic("Assert true failed")
	}
}

func AssertFalse(input bool) {
	if input {
		panic("Assert false failed")
	}
}
