package migrations

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Vehicle struct {
	ID 			uint 		`gorm:"primaryKey"`
	Name 		string 		`gorm:"type:varchar(100);not null"`
	Price 		float64 	`gorm:"type:decimal(10,2);not null"`
	Duration 	string 		`gorm:"type:varchar(10);not null"`
	CreatedAt 	time.Time 	`gorm:"autoCreateTime"`
	UpdatedAt 	time.Time 	`gorm:"autoUpdateTime"`
	IsDeleted 	bool 		`gorm:"default:false"`

	// Relationships
	Pictures 	[]Picture 	`gorm:"foreignKey:VehicleID;"`
	Features 	[]Feature 	`gorm:"foreignKey:VehicleID;"`
}

type Picture struct {
	ID			uint 		`gorm:"primaryKey"`
	VehicleID	uint 		`gorm:"not null;index"`
	URL			string 		`gorm:"type:varchar(255);not null"`
	CreatedAt	time.Time 	`gorm:"autoCreateTime"`
}

type Feature struct {
	ID			uint 		`gorm:"primaryKey"`
	VehicleID	uint 		`gorm:"not null;index"`
	Description	string 		`gorm:"type:varchar(255);not null"`
	CreatedAt	time.Time 	`gorm:"autoCreateTime"`
	UpdatedAt	time.Time 	`gorm:"autoUpdateTime"`
	IsDeleted	bool 		`gorm:"default:false"`
}

type User struct {
	ID			uint 		`gorm:"primaryKey"`
	Name		string 		`gorm:"type:varchar(100);not null"`
	Email		string 		`gorm:"type:varchar(100);unique;not null"`
	Password  	string    	`gorm:"type:varchar(255);not null"`
    Role      	string    	`gorm:"type:varchar(50);not null"`
    CreatedAt 	time.Time 	`gorm:"autoCreateTime"`
    UpdatedAt 	time.Time 	`gorm:"autoUpdateTime"`
    IsDeleted 	bool      	`gorm:"default:false"`
}

func Run() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
  
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_name := os.Getenv("DB_NAME")
	
	dsn := ""+ db_user + ":" + db_password + "@tcp("+ db_host + ":"+ db_port +")/"+ db_name +"?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	err = db.AutoMigrate(&Vehicle{}, &Picture{}, &Feature{}, &User{})
	if err != nil {
		panic("failed to migrate database: " + err.Error())
	}

	println("Database migrated successfully")
}