package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Ligação de muitos para muitos
// Diversos produtos em diversas categorias
// mysql> show tables;
// +-----------------------+
// | Tables_in_productsm2m |
// +-----------------------+
// | categories            |
// | products              |
// | products_categories   |
// +-----------------------+
// 3 rows in set (0.00 sec)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"`
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:products_categories;"`

	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/productsm2m?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{})

	// category_cozinha := Category{Name: "Cozinha"}
	// db.Create(&category_cozinha)

	// category_eletronicos := Category{Name: "Eletronicos"}
	// db.Create(&category_eletronicos)

	// product := Product{
	// 	Name:       "Polo",
	// 	Price:      45.6,
	// 	Categories: []Category{category_cozinha, category_eletronicos},
	// }

	// db.Create(&product)

	var categories []Category
	// Para exibir com o serial number
	// adicionar no Preload Products.SerialNumber. O ProductS é do campo na struct category
	// ele pega tanto o products quanto o serial number
	err = db.Model(&Category{}).
		Preload("Products").
		Find(&categories).
		Error

	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			fmt.Println("- ", product.Name)
		}
	}
}
