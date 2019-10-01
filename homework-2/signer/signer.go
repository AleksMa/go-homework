package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
)

func SingleHash(in, out chan interface{}) {
	for data := range in {
		dataInt := data.(int)
		dataString := strconv.Itoa(dataInt)
		fmt.Printf("SingleHash %s: data %s \n", dataString, dataString)
		md5 := DataSignerMd5(dataString)
		fmt.Printf("SingleHash %s: md5 %s \n", dataString, md5)
		crc32md5 := DataSignerCrc32(md5)
		fmt.Printf("SingleHash %s: crc32md5 %s \n", dataString, crc32md5)
		crc32 := DataSignerCrc32(dataString)
		fmt.Printf("SingleHash %s: crc32 %s \n", dataString, crc32)
		out <- crc32 + "~" + crc32md5
	}
	close(out)
}

func MultiHash(in, out chan interface{}) {
	for data := range in {
		dataString := data.(string)
		result := ""
		for th := 0; th <= 5; th++ {
			crc32 := DataSignerCrc32(string(th) + dataString)
			fmt.Printf("MultiHash %s: step %d: result %s \n", dataString, th, crc32)
			result += crc32
		}
		out <- result
	}
	close(out)
}

func CombineResults(in, out chan interface{}) {
	var dataSlice []string
	for data := range in {
		dataString := data.(string)
		fmt.Printf("CombineResults: got %s \n", dataString)
		dataSlice = append(dataSlice, dataString)
	}
	sort.Strings(dataSlice)
	result := ""
	for i, elem := range dataSlice {
		result += elem
		if i != len(dataSlice)-1 {
			result += "_"
		}
	}
	fmt.Printf("CombineResults: result %s \n", result)
	out <- result
	close(out)
}

func ExecutePipeline(Jobs ...job) {
	in := make(chan interface{})
	out := make(chan interface{})

	for _, Job := range Jobs {
		go Job(in, out)

		in = out
		out = make(chan interface{})
	}
}

func main() {
	// inputData := []int{0, 1, 1, 2, 3, 5, 8}
	inputData := []int{0, 1}

	hashSignJobs := []job{
		job(func(in, out chan interface{}) {
			for _, fibNum := range inputData {
				out <- fibNum
			}
			close(out)
		}),
		job(SingleHash),
		job(MultiHash),
		job(CombineResults),
		job(func(in, out chan interface{}) {
			dataRaw := <-in
			data, ok := dataRaw.(string)
			if !ok {
				log.Fatal("cant convert result data to string")
			}

			fmt.Printf("Result %s \n", data)
		}),
	}

	ExecutePipeline(hashSignJobs...)

	fmt.Scanln()
}
