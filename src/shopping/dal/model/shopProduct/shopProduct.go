package shopProduct

import (
	"shopping/dal/driver"
	"time"
	"shopping/dal/model/product"
	"shopping/dal/model/shoplist"
)

var db = driver.GetDB()

type ShopProduct struct {
	Id           int
	Name         string
	CreationDate time.Time
	Cost         int
	Result       bool
	Product      product.Product
	Photo        string
	Shoplist     shoplist.Shoplist
}
