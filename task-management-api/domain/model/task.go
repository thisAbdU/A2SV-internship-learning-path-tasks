package model

type TaskInfo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
}

type TaskUpdate  struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}