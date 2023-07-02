from locust import HttpUser, task

class LoadTestCourses(HttpUser):
	getAbsenceSessionByClassCode: str = "/api/v1/absence/session/620207"
	getAbsenceByAbsenceSessionId: str = "/api/v1/absence/session-id/d4fc7e23-770f-4f73-a5e1-4d09fd7aa400"
	getClassByNimDosen: str = "/api/v1/siakng/class/nim"
	token: str = "jwt-token"

	@task
	def getabsencesessionbyclasscode(self):
		self.client.get(self.getAbsenceSessionByClassCode, headers={"Authorization": "Bearer {}".format(self.token)})

	@task
	def getabsencebyabsencesessionid(self):
		self.client.get(self.getAbsenceByAbsenceSessionId, headers={"Authorization": "Bearer {}".format(self.token)})

	@task
	def getclassbynimdosen(self):
		self.client.get(self.getClassByNimDosen,  headers={"Authorization": "Bearer {}".format(self.token)})

