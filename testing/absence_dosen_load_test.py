from locust import HttpUser, task

class LoadTestCourses(HttpUser):
	getAbsenceSessionByClassCode: str = "/api/v1/absence/session/620207"
	getAbsenceByAbsenceSessionId: str = "/api/v1/absence/session-id/d4fc7e23-770f-4f73-a5e1-4d09fd7aa400"
	getClassByNimDosen: str = "/api/v1/siakng/class/nim"
	token: str = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1hIjoiRHIuIElyLiBFcmRlZmkgUmFrdW4gTS5TYy4iLCJ1c2VybmFtZSI6ImlpcyIsIm5wbSI6IjEwMDIxMDUxMDI3MTYwOTU5MSIsImp1cnVzYW4iOnsiZmFjdWx0eSI6IklsbXUgS29tcHV0ZXIiLCJzaG9ydEZhY3VsdHkiOiJGYXNpbGtvbSIsIm1ham9yIjoiU2lzdGVtIEluZm9ybWFzaSAoSW5mb3JtYXRpb24gU3lzdGVtKSIsInByb2dyYW0iOiJTMSBQYXJhbGVsIChTMSBQYXJhbGVsKSJ9LCJyb2xlIjoiRG9zZW4iLCJleHAiOjE3MTczMzQ0NTh9.tm1htNgzOnaY48Z9dqfUbR_Lytgz6dYHeOhxt3HQyag"

	@task
	def getabsencesessionbyclasscode(self):
		self.client.get(self.getAbsenceSessionByClassCode, headers={"Authorization": "Bearer {}".format(self.token)})

	@task
	def getabsencebyabsencesessionid(self):
		self.client.get(self.getAbsenceByAbsenceSessionId, headers={"Authorization": "Bearer {}".format(self.token)})

	@task
	def getclassbynimdosen(self):
		self.client.get(self.getClassByNimDosen,  headers={"Authorization": "Bearer {}".format(self.token)})


