package main

import (
	"fmt"
	"sort"
	"strconv"
	"sync"
)

func SingleHash(in, out chan interface{}) {
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	for data := range in {
		dataInt := data.(int)
		dataString := strconv.Itoa(dataInt)

		wg.Add(1)
		go func() {
			defer wg.Done()
			crc32Chan := make(chan string)
			crc32md5Chan := make(chan string)

			go func() {
				fmt.Printf("SingleHash %s: data %s \n", dataString, dataString)
				md5 := DataSignerMd5(dataString)
				fmt.Printf("SingleHash %s: md5 %s \n", dataString, md5)
				mu.Lock()
				crc32md5 := DataSignerCrc32(md5)
				fmt.Printf("SingleHash %s: crc32md5 %s \n", dataString, crc32md5)
				mu.Unlock()
				crc32md5Chan <- crc32md5
			}()

			go func() {
				crc32 := DataSignerCrc32(dataString)
				fmt.Printf("SingleHash %s: crc32 %s \n", dataString, crc32)
				crc32Chan <- crc32
			}()

			result := ""
			result += <-crc32Chan
			result += "~"
			result += <-crc32md5Chan
			fmt.Printf("SingleHash %s: result %s \n", dataString, result)

			out <- result
		}()
	}
	wg.Wait()
}

func MultiHash(in, out chan interface{}) {
	wg := &sync.WaitGroup{}
	for data := range in {
		dataString := data.(string)
		wg.Add(1)
		go func() {
			defer wg.Done()
			result := ""
			hashes := make([]string, 6)
			innerWg := &sync.WaitGroup{}
			for th := 0; th <= 5; th++ {
				innerWg.Add(1)
				go func(th int) {
					defer innerWg.Done()
					crc32 := DataSignerCrc32(strconv.Itoa(th) + dataString)
					fmt.Printf("MultiHash %s: step %d: result %s \n", dataString, th, crc32)
					hashes[th] = crc32
				}(th)
			}
			innerWg.Wait()
			for _, hash := range hashes {
				result += hash
			}
			fmt.Printf("MultiHash %s: summary result %s \n", dataString, result)
			out <- result
		}()
	}
	wg.Wait()
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
}

func ExecutePipeline(Jobs ...job) {
	in := make(chan interface{})
	out := make(chan interface{})

	wg := &sync.WaitGroup{}

	for _, Job := range Jobs {
		wg.Add(1)
		go func(Job job, in, out chan interface{}) {
			defer wg.Done()
			Job(in, out)
			close(out)
		}(Job, in, out)

		in = out
		out = make(chan interface{})
	}

	wg.Wait()
}

func main() {
	// inputData := []int{0, 1, 1, 2, 3, 5, 8}
	inputData := []int{0, 1}

	hashSignJobs := []job{
		job(func(in, out chan interface{}) {
			for _, fibNum := range inputData {
				out <- fibNum
			}
		}),
		job(SingleHash),
		job(MultiHash),
		job(CombineResults),
		job(func(in, out chan interface{}) {
			dataRaw := <-in
			data, ok := dataRaw.(string)
			if !ok {
				fmt.Println("cant convert result data to string")
			}
			fmt.Printf("Result: %s \n", data)
		}),
	}

	ExecutePipeline(hashSignJobs...)
}
