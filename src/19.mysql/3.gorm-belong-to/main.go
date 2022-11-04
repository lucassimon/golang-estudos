package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 1 produto pertence ha varias categorias

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	CategoryID int
	Category   Category
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/products?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{})

	// comentado para nao recriar
	// category := Category{Name: "foo"}
	// db.Create(&category)

	// product := Product{
	// 	Name:       "bar",
	// 	Price:      9.99,
	// 	CategoryID: category.ID,
	// }

	// db.Create(&product)

	var products []Product

	// necessario usar o preload no caso de relacionamentos
	db.Preload("Category").Find(&products)
	for _, product := range products {
		fmt.Println(product.Name, product.Category.Name)
	}
}
