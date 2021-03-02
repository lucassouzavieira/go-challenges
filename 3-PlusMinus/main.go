package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int32) {
	length := len(arr)
	negatives := 0
	positives := 0
	zeros := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] > 0 {
			positives++
		}

		if arr[i] < 0 {
			negatives++
		}

		if arr[i] == 0 {
			zeros++
		}
	}

	pratio := float32(positives) / float32(length)
	nratio := float32(negatives) / float32(length)
	zratio := float32(zeros) / float32(length)

	os.Stdout.WriteString(fmt.Sprint(pratio) + "\n")
	os.Stdout.WriteString(fmt.Sprint(nratio) + "\n")
	os.Stdout.WriteString(fmt.Sprint(zratio) + "\n")
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	plusMinus(arr)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
