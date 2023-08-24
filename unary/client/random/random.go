package random

import (
	"math/rand"
	"time"
)

var (
	productList     = []string{"Oukitel WP27", "OnePlus Ace 2 Pro", "Nokia C210", "Nokia G310"}
	descriptionList = []string{"it's a great phone", "it's a good phone", "it's bad phone", "it's a terrible phone"}
)

func init() {
	rand.Seed(time.Now().Unix())
}
func RandomProduct() string {
	return productList[rand.Intn(len(productList))]
}

func RandomDescription() string {
	return descriptionList[rand.Intn(len(descriptionList))]
}
