package models

type ListCourses struct {
	Courses []CourseMoodle `json:"courses"`
}

type ListQuizzez struct {
	CourseQuizzez []CourseQuizzez `json:"quizzes"`
}

type CourseMoodle struct {
	ID          int                `json:"id"`
	Name        string             `json:"fullname"`
	Assignments []CourseAssignment `json:"assignments,omitempty"`
}

type CourseAssignment struct {
	ID           int64  `json:"cmid"`
	Name         string `json:"name"`
	DueDate      int64  `json:"duedate"`
	TimeModified int64  `json:"timemodified"`
}

type CourseQuizzez struct {
	ID           int64  `json:"coursemodule"`
	Name         string `json:"name"`
	TimeOpen     int64  `json:"timeopen"`
	TimeModified int64  `json:"timemodified"`
	TimeClose    int64  `json:"timeclose"`
}

type Moodle struct {
	User []UserMoodle `json:"users"`
}

type UserMoodle struct {
	ID              int64  `json:"id"`
	Username        string `json:"username"`
	ProfileImageUrl string `json:"profileimageurl"`
}
