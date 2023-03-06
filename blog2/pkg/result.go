package pkg

const (
	Success = 200
	Error   = 500
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

var Msflg = map[int]string{
	Success: "ok",
	Error:   "fail",
}
