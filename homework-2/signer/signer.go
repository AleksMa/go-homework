package main

import "sort"

func SingleHash(in, out chan interface{}) {
	for data := range in {
		dataString := data.(string)
		out <- DataSignerCrc32(dataString) + "~" + DataSignerCrc32(DataSignerMd5(dataString))
	}
}

func MultiHash(in, out chan interface{}) {
	for data := range in {
		dataString := data.(string)
		result := ""
		for i := 0; i <= 5; i++ {
			result += DataSignerCrc32(string(i) + dataString)
		}
		out <- result
	}
}

func CombineResults(in, out chan interface{}) {
	var dataSlice []string
	for data := range in {
		dataSlice = append(dataSlice, data.(string))
	}
	sort.Strings(dataSlice)
	result := ""
	for i, elem := range dataSlice {
		result += elem
		if i != len(dataSlice)-1 {
			result += "_"
		}
	}
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
