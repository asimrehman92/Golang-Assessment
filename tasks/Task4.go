package tasks

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

type Product struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Price uint
}

type Order struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	ProductID uint
}

type Payment struct {
	ID      uint `gorm:"primaryKey"`
	OrderID uint
	Amount  float64
	Status  string
}

var db *gorm.DB

func Task4() {
	r := gin.Default()

	var err error
	db, err = gorm.Open(mysql.Open("user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{}, &Product{}, &Order{}, &Payment{})

	r.POST("/api/transaction", handleTransaction)

	r.Run(":8080")
}

func handleTransaction(c *gin.Context) {
	tx := db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			c.JSON(500, gin.H{"error": "Internal Server Error"})
		}
	}()

	user := User{Name: "John"}
	product := Product{Name: "Laptop", Price: 1500}
	order := Order{UserID: 1, ProductID: 1}
	payment := Payment{OrderID: 1, Amount: 1500, Status: "Success"}

	if err := tx.Create(&user).Error; err != nil {
		panic(err)
	}

	if err := tx.Create(&product).Error; err != nil {
		panic(err)
	}

	if err := tx.Create(&order).Error; err != nil {
		panic(err)
	}

	if err := tx.Create(&payment).Error; err != nil {
		panic(err)
	}

	tx.Commit()

	c.JSON(200, gin.H{"message": "Transaction completed successfully"})
}
