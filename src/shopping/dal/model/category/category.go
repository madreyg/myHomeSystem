package category

import "shopping/dal/driver"

var db = driver.GetDB()

type Category struct {
	Id          int
	Name        string
	Photo       string
	DirectoryId int
}
