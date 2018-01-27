package shoplist

import "shopping/dal/driver"

var db = driver.GetDB()

type Shoplist struct {
	Id   int
	Name string
	Date string
}
