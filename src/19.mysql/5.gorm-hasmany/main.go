package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Ligação de 1 para muitos
// 1 categoria tem muitos produtos

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product
}

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/products?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	var categories []Category
	// Para exibir com o serial number
	// adicionar no Preload Products.SerialNumber. O ProductS é do campo na struct category
	// ele pega tanto o products quanto o serial number
	err = db.Model(&Category{}).
		// Preload("Products").
		Preload("Products.SerialNumber").
		Find(&categories).
		Error

	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			fmt.Println("- ", product.Name, product.Category.Name)
		}
	}

	// Imprime com o SerialNumber
	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			fmt.Println("- ", product.Name, product.Category.Name, product.SerialNumber.Number)
		}
	}
}
