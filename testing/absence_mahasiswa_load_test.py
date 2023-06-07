from locust import HttpUser, task

class LoadTestCourses(HttpUser):
	getClassByNpmMahasiswa: str = "/api/v1/siakng/class/npm"
	checkAbsenceisOpen: str = "/api/v1/absence/check/620207"
	getAbsenceByClassCode: str = "/api/v1/absence/620207"
	doAbsence: str = "/api/v1/absence/"
	token: str = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1hIjoiUmFzeWlkIE1pZnRhaHVsIEloc2FuIiwidXNlcm5hbWUiOiJyYXN5aWQubWlmdGFodWwiLCJucG0iOiIxOTA2MjkzMzAzIiwianVydXNhbiI6eyJmYWN1bHR5IjoiSWxtdSBLb21wdXRlciIsInNob3J0RmFjdWx0eSI6IkZhc2lsa29tIiwibWFqb3IiOiJJbG11IEtvbXB1dGVyIChDb21wdXRlciBTY2llbmNlKSIsInByb2dyYW0iOiJTMSBSZWd1bGVyIChVbmRlcmdyYWR1YXRlIFByb2dyYW0pIn0sInJvbGUiOiJNYWhhc2lzd2EiLCJleHAiOjE4MTIwMDk3NjR9.FWYopWkq-YgSMy72xdNniR2ZBoDxQT_Sg-h-g2M--tY"

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

