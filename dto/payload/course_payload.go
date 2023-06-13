package payload

type ChooseCourse struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type UserEventPayload struct {
	IsDone *bool `json:"is_done"`
}
