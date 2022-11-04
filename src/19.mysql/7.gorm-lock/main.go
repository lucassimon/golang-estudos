package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Lock otimista temos
// name email version
// w    w@w.com 1
// No lock otimista ele verifica a versão novamente e
// refaz todo o processamento novamente caso a versao tenha mudado
// Muitas transações com poucas concorrencias

// Lock pessimista temos
// Locka a tabela naquele registro. Naquele momento ninguem atualiza
// dados
// Usado em grandes aplicações que tem diversas concorrencias

// bloqueia o registro atraves do FOR UPDATE
// 2022/11/04 18:03:11 /home/lucas/go/pkg/mod/gorm.io/gorm@v1.24.1/callbacks.go:134
// [0.360ms] [rows:1] SELECT * FROM `categories` WHERE `categories`.`id` = 1 ORDER BY `categories`.`id` LIMIT 1 FOR UPDATE

// 2022/11/04 18:03:11 /home/lucas/go/pkg/mod/gorm.io/gorm@v1.24.1/callbacks.go:134
// [0.340ms] [rows:1] UPDATE `categories` SET `name`='Foo Bar' WHERE `id` = 1

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
	dsn := "root:root@tcp(localhost:3306)/productslock?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{})

	// category_eletronicos := Category{Name: "Eletronicos"}
	// db.Create(&category_eletronicos)

	// inicia a transaction
	transaction := db.Begin()
	var category Category
	err = transaction.
		Debug().
		Clauses(clause.Locking{Strength: "UPDATE"}).
		First(&category, 1).Error

	if err != nil {
		panic(err)
	}
	category.Name = "Foo Bar"
	transaction.Debug().Save(&category)
	transaction.Commit()
}
