package task

type Task struct {
    ID          string `json:"id" bson:"_id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Status      string `json:"status"`
}
