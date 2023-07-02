from locust import HttpUser, task

class LoadTestCourses(HttpUser):
	getClassByNpmMahasiswa: str = "/api/v1/siakng/class/npm"
	checkAbsenceisOpen: str = "/api/v1/absence/check/620207"
	getAbsenceByClassCode: str = "/api/v1/absence/620207"
	doAbsence: str = "/api/v1/absence/"
	token: str = "jwt-token"

	@task
	def getclassbynpmmahasiswa(self):
		self.client.get(self.getClassByNpmMahasiswa, headers={"Authorization": "Bearer {}".format(self.token)})

	@task
	def checkabsenceisopen(self):
		self.client.get(self.checkAbsenceisOpen, headers={"Authorization": "Bearer {}".format(self.token)})

	@task
	def getabsencebyclasscode(self):
		self.client.get(self.getAbsenceByClassCode,  headers={"Authorization": "Bearer {}".format(self.token)})

	@task
	def doabsence(self):
		self.client.put(self.doAbsence, headers={"Authorization": "Bearer {}".format(self.token)}, json={
			"class_session_id" : "d4fc7e23-770f-4f73-a5e1-4d09fd7aa400",
    		"device_code" : "Iphone 12 Pro"
		})

