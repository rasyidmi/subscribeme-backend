package utils

type DataNotFound struct {
}

type DifferentSubjectID struct {
}

type TokenNotAvailable struct {
}

func (DataNotFound) Error() string {
	return "Data is not found."
}

func (DifferentSubjectID) Error() string {
	return "New data subject_id is different with the old one."
}

func (TokenNotAvailable) Error() string {
	return "Token is not exist."
}
