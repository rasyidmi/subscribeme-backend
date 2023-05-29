package models

type ClassSchedule struct {
	ScheduleUrl string      `json:"url"`
	Day         string      `json:"hari"`
	StartTime   string      `json:"jam_mulai"`
	EndTime     string      `json:"jam_selesai"`
	ClassDetail ClassDetail `json:"kd_kls_sc"`
	Room        Room        `json:"id_ruang"`
}

type ClassDetail struct {
	ClassName   string        `json:"nm_kls"`
	ClassCode   string        `json:"kd_kls"`
	Course      Course        `json:"nm_mk_cl,omitempty"`
	Lecturers   []Lecturers   `json:"pengajar"`
	ListStudent []ListStudent `json:"daftar_mahasiswa,omitempty"`
}

type Course struct {
	CourseCode string `json:"kd_mk"`
	CourseName string `json:"nm_mk"`
	CourseSKS  int    `json:"jml_sks"`
}

type Lecturers struct {
	Name string `json:"nama"`
}

type Room struct {
	RoomName string `json:"nm_ruang"`
}

type ListStudent struct {
	Student Student `json:"mahasiswa"`
}

type Student struct {
	Name string `json:"nama"`
	Npm  string `json:"npm"`
}
