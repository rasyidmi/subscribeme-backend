from locust import HttpUser, task

class LoadTestCourses(HttpUser):
	getCoursesSceleByUsername: str = "/api/v1/moodle/courses/username"
	getEnrolledCoursesByUsername: str = "/api/v1/course"
	getUserEventByCourseId: str = "/api/v1/course/event/0d038ce7-191c-44a9-a694-fdb534981db3"
	getTodayDeadline: str = "/api/v1/course/deadline/today"
	getSevenAheadDeadline: str = "/api/v1/course/deadline/7-days"
	subscribeCourse: str = "/api/v1/course/subscribe"
	unsubscribeCourse: str = "/api/v1/course/unsubscribe"
	token: str = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1hIjoiUmFzeWlkIE1pZnRhaHVsIEloc2FuIiwidXNlcm5hbWUiOiJyYXN5aWQubWlmdGFodWwiLCJucG0iOiIxOTA2MjkzMzAzIiwianVydXNhbiI6eyJmYWN1bHR5IjoiSWxtdSBLb21wdXRlciIsInNob3J0RmFjdWx0eSI6IkZhc2lsa29tIiwibWFqb3IiOiJJbG11IEtvbXB1dGVyIChDb21wdXRlciBTY2llbmNlKSIsInByb2dyYW0iOiJTMSBSZWd1bGVyIChVbmRlcmdyYWR1YXRlIFByb2dyYW0pIn0sInJvbGUiOiJNYWhhc2lzd2EiLCJleHAiOjE4MTIwMDk3NjR9.FWYopWkq-YgSMy72xdNniR2ZBoDxQT_Sg-h-g2M--tY"

	@task
	def getcoursesscelebyusername(self):
		self.client.get(self.getCoursesSceleByUsername, headers={"Authorization": "Bearer {}".format(self.token)})

	@task
	def getenrolledcoursesbyusername(self):
		self.client.get(self.getEnrolledCoursesByUsername, headers={"Authorization": "Bearer {}".format(self.token)})

	@task
	def getusereventbycourseid(self):
		self.client.get(self.getUserEventByCourseId,  headers={"Authorization": "Bearer {}".format(self.token)})

	@task
	def gettodaydeadline(self):
		self.client.get(self.getTodayDeadline,  headers={"Authorization": "Bearer {}".format(self.token)})

	@task
	def getsevenaheaddeadline(self):
		self.client.get(self.getSevenAheadDeadline,  headers={"Authorization": "Bearer {}".format(self.token)})

	@task
	def subscribecourse(self):
		self.client.post(self.subscribeCourse, headers={"Authorization": "Bearer {}".format(self.token)}, json= {
			"id": 4,
    		"name": "Matematika Diskrit"
		})
	@task
	def unsubscribecourse(self):
		self.client.post(self.unsubscribeCourse, headers={"Authorization": "Bearer {}".format(self.token)}, json= {
			"id": 4,
    		"name": "Matematika Diskrit"
		})
