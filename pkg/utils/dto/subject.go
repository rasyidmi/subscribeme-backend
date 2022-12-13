package dto

type SubjectRequest struct {
	ID      int
	Title   string              `json:"title"`
	Term    int                 `json:"term"`
	Major   string              `json:"major"`
	Classes []map[string]string `json:"classes"`
}

type SubjectResponse struct {
	ID      int
	Title   string
	Term    int
	Major   string
	Classes []map[string]string
}
