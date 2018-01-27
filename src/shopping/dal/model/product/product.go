package product

import (
	"shopping/dal/driver"
	"log"
	"errors"
)

var db = driver.GetDB()

type Product struct {
	Id         int
	Name       string
	Photo      string
	CategoryId int
}

var FieldAll = `
				product.id,
				product.name,
				product.photo,
				product.category_id
			   `

func New(d *Product) (error) {
	log.Println("New product.")
	_, err := db.Exec(` INSERT INTO product (name, photo) VALUES ($1::TEXT, $2::TEXT, $3::INT) `,
		d.Name, d.Photo, d.CategoryId)
	if err != nil {
		return err
	}
	log.Println("New product complecte.")
	return nil
}

func (d *Product) Update() (error) {
	log.Println("Update product.")
	query := "UPDATE product SET name = $1::text, photo = $2::text, category_id = $3::int " +
		"WHERE id = $4::int"
	res, err := db.Exec(query, d.Name, d.Photo, d.CategoryId, d.Id)
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
	log.Println("Update product complete")
	return nil
}

func (d *Product) Delete() (error) {
	log.Println("Delete product ")
	_, err := db.Exec("DELETE FROM product WHERE id = $1:int", d.Id)
	if err != nil {
		return err
	}
	log.Println("Delete product comlete")
	return nil
}

func GetAll() ([]Product, error) {
	query := "SELECT " + FieldAll + "FROM product"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	products := *new([]Product)
	productTmp := new(Product)
	for rows.Next() {
		err := rows.Scan(&productTmp.Id, &productTmp.Name, &productTmp.Photo)
		if err != nil {
			return nil, err
		}
		products = append(products, *productTmp)
	}
	return products, nil
}
