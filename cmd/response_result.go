package cmd

import (
	"fmt"
	"time"
)

type ResponseResultList []ResponseResult

type ResponseResult struct {
	Url       string        `json:"url"`
	Method    string        `json:"method"`
	Status    int           `json:"status"`
	TimeTaken time.Duration `json:"time_taken"`
}

func (r ResponseResultList) RenderOutput(t time.Duration) {
	var sum, avg, min, max time.Duration
	var success, fail int
	for i, v := range r {
		if v.Status < 400 {
			success++
		} else {
			fail++
		}
		sum += v.TimeTaken
		if i == 0 {
			max, min = v.TimeTaken, v.TimeTaken
			continue
		}
		if max < v.TimeTaken {
			max = v.TimeTaken
		}
		if min > v.TimeTaken {
			min = v.TimeTaken
		}
	}
	avg = time.Duration(sum.Nanoseconds() / int64(len(r)))
	fmt.Println("success: ", success)
	fmt.Println("fail: ", fail)
	fmt.Println("exec_time: ", t)
	fmt.Println("avg: ", avg)
	fmt.Println("min: ", min)
	fmt.Println("max: ", max)
}
