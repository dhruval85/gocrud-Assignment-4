package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

type User struct {
	gorm.Model
	ID      uint   `gorm:"primaryKey"`
	Name    string `gorm:"not null"`
	Address string `gorm:"not null"`
}

func main() {
	connectDB()
	// Migrate the schema
	err = DB.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	//createUser()
	//createMultipleUsers()
	//getAllUsers()
	//updateUser(2)
	//deleteUser(1)

}

func connectDB() {
	dsn := "root:ADMIN@tcp(127.0.0.1:3306)/gotest_db?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	log.Println("Database connection successful!")

}

// Function to create  users
func createUser() {
	user := User{ID: 5, Name: "Arshdeep", Address: "Punjab"}
	DB.Create(&user)
	log.Println("New user created:", user.ID)
}

// Function to create multiple users
func createMultipleUsers() {
	users := []User{
		{Name: "Arshdeep", Address: "Punjab"},
		{Name: "Rohit Sharma", Address: "Mumbai"},
		{Name: "Virat Kohli", Address: "Delhi"},
	}

	result := DB.Create(&users) // Insert multiple users

	if result.Error != nil {
		log.Fatalf("failed to create users: %v", result.Error)
	}

	for _, user := range users {
		log.Println("New user created:", user.ID)
	}
}

// Function to fetch users
func getAllUsers() {
	var users []User
	DB.Find(&users)
	log.Println("All users:", users)
}

// Function to update  users
func updateUser(id uint) {
	var user User
	DB.First(&user, id)
	user.Name = "Hardik pandya "
	user.Address = " Gujarat"
	DB.Save(&user)
	log.Println("User updated:", user.ID)
}

// Function to delete users
func deleteUser(id uint) {
	var user User
	DB.Delete(&user, id)
	log.Println("User deleted:", user.ID)
}
