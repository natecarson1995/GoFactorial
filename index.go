package main

import (
	"fmt"
	"os"
	"math/big"
	"strconv"
)

func multiplyList(resultChannel chan<- big.Int, numberList []int) {
	switch len := len(numberList); {
	case len <= 1:
		resultChannel <- *big.NewInt( int64(numberList[0]) )

	case len <= 2:
		product, multiplier := big.NewInt( int64(numberList[0]) ), big.NewInt( int64(numberList[1]) )

		product.Mul(product, multiplier)
		resultChannel <- *product
	
	default:
		subchannel := make(chan big.Int)

		go multiplyList(subchannel, numberList[:len/2])
		go multiplyList(subchannel, numberList[len/2:])
		
		product, multiplier := <-subchannel, <-subchannel
		product.Mul(&product, &multiplier)

		resultChannel <- product
	}
}

func makeRange(min, max int) []int {
    rangeList := make([]int, max-min+1)
    for i := range rangeList {
        rangeList[i] = min + i
    }
    return rangeList
}

func factorial(i int) big.Int {
	channel, rangeList := make(chan big.Int), makeRange(1, i)

	go multiplyList(channel, rangeList)
	return <-channel
}

func main() {
	arg := 10
	if len(os.Args) >= 1 {
		arg, _ = strconv.Atoi(os.Args[1])
	}

	k := factorial(arg)
	fmt.Printf("%v", k.String())

}