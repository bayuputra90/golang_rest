package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	cfg "project/configs"
	m "project/models"

	"github.com/gorilla/mux"
)

func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	// menyiapkan variabel untuk menampung struct Student
	var students m.Students
	// menyiapkan slice untuk Student
	var arr_student []m.Students
	// panggil juga struct Response
	var response m.AllResponse

	db := cfg.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM students")

	if err != nil {
		log.Print(err)
		return
	}

	// var simpanData = []any{
	// 	&students.Id,
	// 	&students.NamaDepan,
	// 	&students.NamaBelakang,
	// 	&students.NoHp,
	// 	&students.Gender,
	// 	&students.Jenjang,
	// 	&students.Hobi,
	// 	&students.Alamat,
	// }

	for rows.Next() {
		err := rows.Scan(&students.Id, &students.NamaDepan, &students.NamaBelakang, &students.NoHp, &students.Gender, &students.Jenjang, &students.Hobi, &students.Alamat)

		if err != nil {
			log.Fatal(err.Error())
		} else {
			arr_student = append(arr_student, students)
		}
	}

	response.Status = 1
	response.Message = "Success"
	response.Data = arr_student

	w.Header().Set("Content-Type", "application/json")
	
	// mengubah data dari database menjadi json karena di encode
	json.NewEncoder(w).Encode(response)

}

func InsertStudent(w http.ResponseWriter, r *http.Request) {

	var response m.AllResponse

	db := cfg.Connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)

	if err != nil {
		panic(err)
	}

	// membaca data dari input form postman / app mobile
	nama_depan := r.FormValue("nama_depan")
	nama_belakang := r.FormValue("nama_belakang")
	no_hp := r.FormValue("no_hp")
	gender := r.FormValue("gender")
	jenjang := r.FormValue("jenjang")
	hobi := r.FormValue("hobi")
	alamat := r.FormValue("alamat")

	_, err = db.Exec("INSERT INTO students (nama_depan, nama_belakang, no_hp, gender, jenjang, hobi, alamat) values (?,?,?,?,?,?,?)",
		nama_depan,
		nama_belakang,
		no_hp,
		gender,
		jenjang,
		hobi,
		alamat,
	)

	if err != nil {
		log.Print(err)
	}

	response.Status = 1
	response.Message = "Berhasil insert data siswa"
	log.Print("Berhasil insert data siswa")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func GetStudent(w http.ResponseWriter, r *http.Request) {
	var student m.Students
	var response m.Response

	db := cfg.Connect()
	defer db.Close()

	// id := r.FormValue("id")

	// ambil data by parameter
	param := mux.Vars(r)
	id := param["id"]

	rows, err := db.Query("Select * FROM students WHERE id = ?",
		id,
	)

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {

		err := rows.Scan(&student.Id, &student.NamaDepan, &student.NamaBelakang, &student.NoHp, &student.Gender, &student.Jenjang, &student.Hobi, &student.Alamat)

		if err != nil {
			log.Fatal(err.Error())

		}

	}

	response.Status = 1
	response.Message = "Success"
	response.Data = student

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {

	var response m.Responses

	db := cfg.Connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)

	if err != nil {
		panic(err)
	}

	param := mux.Vars(r)
	// id := r.FormValue("id")
	id := param["id"]

	nama_depan := r.FormValue("nama_depan")
	nama_belakang := r.FormValue("nama_belakang")
	no_hp := r.FormValue("no_hp")
	gender := r.FormValue("gender")
	jenjang := r.FormValue("jenjang")
	hobi := r.FormValue("hobi")
	alamat := r.FormValue("alamat")

	_, err = db.Exec("UPDATE students SET nama_depan = ?, nama_belakang = ?, no_hp = ?, gender = ?, jenjang = ?, hobi = ?, alamat = ? WHERE id = ?",
		nama_depan,
		nama_belakang,
		no_hp,
		gender,
		jenjang,
		hobi,
		alamat,
		id,
	)

	if err != nil {
		log.Print(err)
	}

	response.Status = 1
	response.Message = "Berhasil Update Data Siswa"
	log.Print("Berhasil Update Data Siswa")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {

	var response m.Responses
	db := cfg.Connect()
	defer db.Close()

	param := mux.Vars(r)
	id := param["id"]

	_, err := db.Exec("DELETE FROM students WHERE id = ?",
		id,
	)

	if err != nil {
		log.Print(err)
	}

	response.Status = 1
	response.Message = "Berhasil Delete Data Siswa"
	log.Print("Berhasil Delete Data Siswa")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
