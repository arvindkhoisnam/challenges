package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


type DbConfig struct{
	Host string 
	User string
	Password string
	DbName string
	Port string
	SslMode string
}

func GenerateClient(config *DbConfig)(*gorm.DB,error){
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",config.Host,config.User,config.Password,config.DbName,config.Port,config.SslMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	if err != nil{
		return nil ,err
	}
	return db,nil
}
//retry logic
// for i := 0; i < 3;i ++{
// 	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",config.Host,config.User,config.Password,config.DbName,config.Port,config.SslMode)
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err == nil{
// 		return db
// 	}
// 	fmt.Println("Database not ready, retrying in 2s...")
// 	time.Sleep(2 * time.Second)
// }
// panic("Could not connect to database")