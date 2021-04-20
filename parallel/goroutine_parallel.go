package parallel

func StreamOf(nums []int, out chan <- int) {
	for _, v := range nums {
		out <- v
	}
	close(out)
}


func MapToStr(inChan <- chan int, outChan chan <- string, mapFunc func(num int) string) {
	for i := range inChan {
		str := mapFunc(i)
		outChan <- str
	}
	close(outChan)
}