package constant

type UserRoleEnum string

const (
	UserRoleMahasiswa UserRoleEnum = "Mahasiswa"
	UserRoleDosen     UserRoleEnum = "Dosen"
)

func (x UserRoleEnum) String() string {
	return string(x)
}
