package main

import (
	"encoding/json"
	"fmt"
)

func getDoctorsFromDB() []byte {
	var (
		doctor  Doctor
		doctors []Doctor
	)

	rows, err := db.Query("SELECT doctorsID, doctorFirstName, doctorLastName, doctorDegree, doctorReg, doctorOnline, doctorsWherebyLink, doctorsPicture  FROM doctors")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	//the sequence of fields matters. the sequence of fields below must be same as above
	//in rows
	for rows.Next() {
		rows.Scan(&doctor.DoctorsID, &doctor.DoctorFirstName, &doctor.DoctorLastName,
			&doctor.DoctorDegree, &doctor.DoctorReg, &doctor.DoctorOnline, &doctor.DoctorsWherebyLink, &doctor.DoctorsPicture)
		doctors = append(doctors, doctor)
	}
	defer rows.Close()
	jsonResponse, jsonError := json.Marshal(doctors)
	if jsonError != nil {
		fmt.Println(jsonError)
		return nil
	}
	return jsonResponse
}

/*
func insertUserInDB(userDetails User) bool {
	stmt, err := db.Prepare("INSERT into test.users SET id=?, Name=?,Lname=?,Country=?")
	if err != nil {
		fmt.Println(err)
		return false
	}
	_, queryError := stmt.Exec(userDetails.ID, userDetails.Name, userDetails.Lname, userDetails.Country)
	if queryError != nil {
		fmt.Println(queryError)
		return false
	}
	return true
}

func deleteUserFromDB(userID string) bool {
	stmt, err := db.Prepare("DELETE FROM test.users WHERE id=?")
	if err != nil {
		fmt.Println(err)
		return false
	}
	_, queryError := stmt.Exec(userID)
	if queryError != nil {
		fmt.Println(queryError)
		return false
	}
	return true
}

func updateUserInDB(userDetails User) bool {
	stmt, err := db.Prepare("UPDATE test.users SET name=?,lname=?,country=? WHERE id=?")
	if err != nil {
		fmt.Println(err)
		return false
	}
	_, queryError := stmt.Exec(userDetails.Name, userDetails.Lname, userDetails.Country, userDetails.ID)
	if queryError != nil {
		fmt.Println(queryError)
		return false
	}
	return true
}

*/
