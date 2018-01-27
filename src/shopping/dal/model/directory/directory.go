package directory

import (
	"shopping/dal/driver"
	"log"
	"errors"
)

var db = driver.GetDB()

type Directory struct {
	Id    int
	Name  string
	Photo string
}

var FieldAll = `
				directory.id,
				directory.name,
				directory.photo
			   `

func New(d *Directory) (error) {
	log.Println("New directory.")
	_, err := db.Exec(` INSERT INTO directory (name, photo) VALUES ($1::TEXT, $2::TEXT) `, d.Name, d.Photo)
	if err != nil {
		return err
	}
	log.Println("New directory complecte.")
	return nil
}

func (d *Directory) Update() (error) {
	log.Println("Update directory.")
	query := "UPDATE directory SET name = $1::text, photo = $2::text WHERE id = $3::int"
	res, err := db.Exec(query, d.Name, d.Photo, d.Id)
	if err != nil {
		return err
	}
	aff, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if aff == 0 {
		return errors.New("nothing update")
	}
	log.Println("Update directory complete")
	return nil
}

func (d *Directory) Delete() (error) {
	log.Println("Delete directory ")
	_, err := db.Exec("DELETE FROM directory WHERE id = $1:int", d.Id)
	if err != nil {
		return err
	}
	log.Println("Delete directory comlete")
	return nil
}

func GetAll() ([]Directory, error) {
	query := "SELECT " + FieldAll + "FROM directory"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	directyries := *new([]Directory)
	directoryTmp := new(Directory)
	for rows.Next() {
		err := rows.Scan(&directoryTmp.Id, &directoryTmp.Name, &directoryTmp.Photo)
		if err != nil {
			return nil, err
		}
		directyries = append(directyries, *directoryTmp)
	}
	return directyries, nil
}

