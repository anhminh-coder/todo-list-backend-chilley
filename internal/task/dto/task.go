package dto

type CreateTaskReq struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateTaskReq struct {
	Completed bool `json:"completed"`
}
