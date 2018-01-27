package directory

import "shopping/dal/driver"

var db = driver.GetDB()

type Directory struct {
	Id    int
	Name  string
	Photo string
}
