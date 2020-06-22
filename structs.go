package main

/*
rows, err := db.Query("SELECT
doctorsID,
doctorFirstName,
doctorLastName,
doctorDegree,
doctorReg,
doctorOnline,
doctorsWherebyLink
FROM doctors")

*/
// User is Interface for user details.
type Doctor struct {
	DoctorsID          int    `json:"doctors_id"`
	DoctorFirstName    string `json:"doctor_first_name"`
	DoctorLastName     string `json:"doctor_last_name"`
	DoctorDegree       string `json:"doctor_degree"`
	DoctorReg          string `json:"doctor_reg"`
	DoctorOnline       int    `json:"doctor_online"`
	DoctorsWherebyLink string `json:"doctors_whereby_link"`
	DoctorsPicture     string `json:"doctors_picture"`
}

// ErrorResponse is interface for sending error message with code.
type ErrorResponse struct {
	Code    int
	Message string
}
