package models

import (
	"arnov17/echo-test/db"
	"net/http"

	validator "github.com/go-playground/validator/v10"
)

type Pegawai struct {
	Id      int    `json:"id"`
	Name    string `json:"name" validate:"required"`
	Alamat  string `json:"alamat" validate:"required"`
	Telepon string `json:"telepon" validate:"required"`
}

func FetcAllPegawai() (Response, error) {
	var obj Pegawai
	var arrayObj []Pegawai
	var res Response

	con := db.CreateCont()

	sqlStatemnet := "SELECT * FROM pegawai"

	rows, err := con.Query(sqlStatemnet)

	defer con.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Name, &obj.Alamat, &obj.Telepon)
		if err != nil {
			return res, err
		}

		arrayObj = append(arrayObj, obj)

	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrayObj

	return res, err
}

func StorePegawai(name string, alamat string, telepon string) (Response, error) {
	var res Response

	v := validator.New()

	peg := Pegawai{
		Name:    name,
		Alamat:  alamat,
		Telepon: telepon,
	}

	err := v.Struct(peg)
	if err != nil {
		return res, err
	}

	con := db.CreateCont()
	sqlStatemnet := "INSERT pegawai (nama, alamat, telepon) VALUES (?, ?, ?)"

	stmt, err := con.Prepare(sqlStatemnet)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, alamat, telepon)
	if err != nil {
		return res, err
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Sucess"
	res.Data = map[string]int64{
		"last_insereted_id": lastId,
	}

	return res, nil
}

func UpdatePegawai(id int, nama string, alamat string, telepon string) (Response, error) {
	var res Response

	con := db.CreateCont()

	sqlStatemnet := "UPDATE pegawai SET nama = ?, alamat = ?, telepon = ? WHERE id = ?"

	stmt, err := con.Prepare(sqlStatemnet)
	if err != nil {
		return res, err
	}
	result, err := stmt.Exec(nama, alamat, telepon, id)
	if err != nil {
		return res, err
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affctd": rowsAff,
	}
	return res, err
}

func DeletePegawai(id int) (Response, error) {
	var res Response

	con := db.CreateCont()

	sqlStatement := "DELETE FROM pegawai WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}
