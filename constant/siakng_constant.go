package constant

type SiakNGEnum string

const (
	GetClassScheduleByYearAndTerm  SiakNGEnum = "https://api-kp.cs.ui.ac.id/siakngcs/jadwal-list/%s/%s/"
	GetClassScheduleByScheduleId   SiakNGEnum = "https://api-kp.cs.ui.ac.id/siakngcs/jadwal/%s/"
	GetClassScheduleByNpmMahasiswa SiakNGEnum = "https://api-kp.cs.ui.ac.id/siakngcs/mahasiswa/%s/jadwal-kelas/%s/%s/"
	GetClassParticipantByClassCode SiakNGEnum = "https://api-kp.cs.ui.ac.id/siakngcs/kelas/peserta/%s/"
)

func (x SiakNGEnum) String() string {
	return string(x)
}
