package constant

type SiakNGEnum string

const (
	GetClassDetailByNpmMahasiswa   SiakNGEnum = "https://api-kp.cs.ui.ac.id/siakngcs/mahasiswa/%s/jadwal-kelas/%s/%s/"
	GetClassParticipantByClassCode SiakNGEnum = "https://api-kp.cs.ui.ac.id/siakngcs/kelas/peserta/%s/"

	GetClassDetailByNimDosen SiakNGEnum = "https://api-kp.cs.ui.ac.id/siakngcs/dosen/%s/jadwal-kelas/%s/%s/"
	GetClassByCode           SiakNGEnum = "https://api-kp.cs.ui.ac.id/siakngcs/kelas/%s/"
)

func (x SiakNGEnum) String() string {
	return string(x)
}
