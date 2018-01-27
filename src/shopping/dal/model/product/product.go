package product

import "shopping/dal/driver"

var db = driver.GetDB()

type Product struct {
	Id         int
	Name       string
	Photo      string
	CategoryId int
}

