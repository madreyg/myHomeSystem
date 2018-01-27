package category

import (
	"shopping/dal/driver"
	"log"
	"errors"
)

var db = driver.GetDB()

type Category struct {
	Id          int
	Name        string
	Photo       string
	DirectoryId int
}

var FieldAll = `
				category.id,
				category.name,
				category.photo,
				category.directory_id
			   `

func New(d *Category) (error) {
	log.Println("New category.")
	_, err := db.Exec(` INSERT INTO category (name, photo, directory_id) VALUES ($1::TEXT, $2::TEXT, $3::INT) `,
		d.Name, d.Photo, d.DirectoryId)
	if err != nil {
		return err
	}
	log.Println("New category complecte.")
	return nil
}

func (d *Category) Update() (error) {
	log.Println("Update category.")
	query := "UPDATE category SET name = $1::text, photo = $2::text, directory_id = $3::int " +
		"WHERE id = $4::int"
	res, err := db.Exec(query, d.Name, d.Photo, d.DirectoryId, d.Id)
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
	log.Println("Update category complete")
	return nil
}

func (d *Category) Delete() (error) {
	log.Println("Delete category ")
	_, err := db.Exec("DELETE FROM category WHERE id = $1:int", d.Id)
	if err != nil {
		return err
	}
	log.Println("Delete category comlete")
	return nil
}

func GetAll() ([]Category, error) {
	query := "SELECT " + FieldAll + "FROM category"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	categories := *new([]Category)
	categoryTmp := new(Category)
	for rows.Next() {
		err := rows.Scan(&categoryTmp.Id, &categoryTmp.Name, &categoryTmp.Photo)
		if err != nil {
			return nil, err
		}
		categories = append(categories, *categoryTmp)
	}
	return categories, nil
}
