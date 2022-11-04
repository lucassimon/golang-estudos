package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// gorm.Model aplica campos como created_at updated_at deleted_at
// utilizado para o soft-delete
// mysql> desc products;
// +------------+-------------+------+-----+---------+----------------+
// | Field      | Type        | Null | Key | Default | Extra          |
// +------------+-------------+------+-----+---------+----------------+
// | id         | bigint(20)  | NO   | PRI | NULL    | auto_increment |
// | name       | longtext    | YES  |     | NULL    |                |
// | price      | double      | YES  |     | NULL    |                |
// | created_at | datetime(3) | YES  |     | NULL    |                |
// | updated_at | datetime(3) | YES  |     | NULL    |                |
// | deleted_at | datetime(3) | YES  | MUL | NULL    |                |
// +------------+-------------+------+-----+---------+----------------+
// 6 rows in set (0.00 sec)
type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
	gorm.Model
}

func create(db *gorm.DB) {
	fmt.Println("create")
	db.Create(&Product{
		Name:  "Placa mae",
		Price: 1000.00,
	})

	fmt.Println("create batch")
	products := []Product{
		{Name: "Notebook", Price: 1000.00},
		{Name: "Mouse", Price: 50.00},
		{Name: "Keyboard", Price: 100.00},
	}
	db.Create(&products)
	fmt.Println("---")
}

func selectOne(db *gorm.DB) {
	fmt.Println("select one")
	var product Product
	db.First(&product, 2)
	fmt.Println(product)
	fmt.Println("select one name mouse")
	db.First(&product, "name = ?", "Mouse")
	fmt.Println(product)
	fmt.Println("---")
}

func selectAll(db *gorm.DB) {
	fmt.Println("select all")
	var products []Product
	db.Find(&products)
	for _, product := range products {
		fmt.Println(product)
	}
	fmt.Println("---")
}

func selectAllPaginate(db *gorm.DB) {
	fmt.Println("Limit e offset")
	var products []Product
	db.Limit(2).Offset(2).Find(&products)
	for _, product := range products {
		fmt.Println(product)
	}
	fmt.Println("---")
}

func selectWhere(db *gorm.DB) {
	fmt.Println("where")
	var products []Product
	db.Where("price > ?", 100).Find(&products)
	for _, product := range products {
		fmt.Println(product)
	}
	fmt.Println("---")
}

func selectLike(db *gorm.DB) {
	fmt.Println("Like")
	var products []Product
	db.Where("name LIKE ?", "%key%").Find(&products)
	for _, product := range products {
		fmt.Println(product)
	}
	fmt.Println("---")
}

// mysql> select * from products;
// +----+-----------+-------+-------------------------+-------------------------+------------+
// | id | name      | price | created_at              | updated_at              | deleted_at |
// +----+-----------+-------+-------------------------+-------------------------+------------+
// |  1 | New Mouse |  1000 | 2022-11-03 19:14:28.805 | 2022-11-04 17:07:23.601 | NULL       |
func update(db *gorm.DB) {
	fmt.Println("update")
	var p Product
	db.First(&p, 1)
	p.Name = "New Mouse"
	db.Save(&p)
	fmt.Println("---")
}

// mysql> select * from products;
// +----+-----------+-------+-------------------------+-------------------------+-------------------------+
// | id | name      | price | created_at              | updated_at              | deleted_at              |
// +----+-----------+-------+-------------------------+-------------------------+-------------------------+
// |  1 | New Mouse |  1000 | 2022-11-03 19:14:28.805 | 2022-11-04 17:07:23.601 | 2022-11-04 17:08:26.109 |
func delete(db *gorm.DB) {
	var p2 Product
	db.First(&p2, 1)
	fmt.Println(p2.Name)
	db.Delete(&p2)
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/products?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	// create(db)
	// selectOne(db)
	// selectAll(db)
	// selectAllPaginate(db)
	// selectWhere(db)
	// selectLike(db)
	// update(db)
	// delete(db)
}
