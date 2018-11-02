package cmd

type ResponseResult struct {
	Url       string `json:"url"`
	Method    string `json:"method"`
	Status    string `json:"status"`
	TimeTaken string `json:"time_taken"`
}
