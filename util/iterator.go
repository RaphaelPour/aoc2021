package util

func EachDigit(input, base int) <-chan int {
	iteratorChannel := make(chan int)
	go func(input, base int) {
		for input > 0 {
			iteratorChannel <- input % base
			input /= base
		}
		close(iteratorChannel)
	}(input, base)
	return iteratorChannel
}
