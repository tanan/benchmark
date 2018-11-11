package cmd

import (
	"time"
)

type ResponseResultList []ResponseResult

type ResponseResult struct {
	Url       string        `json:"url"`
	Method    string        `json:"method"`
	Status    string        `json:"status"`
	TimeTaken time.Duration `json:"time_taken"`
}

func (r ResponseResultList) RenderOutput() string {
	var sum, max, min time.Duration
	for _, v := range r {
		sum += v.TimeTaken
		if max > v.TimeTaken {
			max = v.TimeTaken
		}
		if min < v.TimeTaken {
			min = v.TimeTaken
		}
	}
	avg := sum.Nanoseconds() / int64(len(r))
	return string(avg)
}
