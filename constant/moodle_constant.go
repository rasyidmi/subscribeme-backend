package constant

type MoodleEnum string

const (
	GetUserDetailByUsername MoodleEnum = "http://103.13.207.85/webservice/rest/server.php?wstoken=%s&wsfunction=core_user_get_users&moodlewsrestformat=json&criteria[0][key]=username&criteria[0][value]=%s"
	GetCourseByUserid       MoodleEnum = "http://103.13.207.85/webservice/rest/server.php?wstoken=%s&wsfunction=core_enrol_get_users_courses&moodlewsrestformat=json&userid=%d"

	GetAssignmentFromCourseID MoodleEnum = "http://103.13.207.85/webservice/rest/server.php?wstoken=%s&wsfunction=mod_assign_get_assignments&moodlewsrestformat=json&%s=%s"
	GetQuizFromCourseID       MoodleEnum = "http://103.13.207.85/webservice/rest/server.php?wstoken=%s&wsfunction=mod_quiz_get_quizzes_by_courses&moodlewsrestformat=json&%s=%s"
)

func (x MoodleEnum) String() string {
	return string(x)
}
