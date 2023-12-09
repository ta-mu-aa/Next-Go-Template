package db

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbEnv struct {
	Host     string
	User     string
	Pass     string
	DB       string
	Port     string
	TimeZone string
}

var db *gorm.DB
const retry = 10 // 再接続回数
var count = 0     // 接続回数

func Init() (string, error) {
	var env DbEnv

	// 環境変数取得
	env.Host = os.Getenv("HOST") // host=コンテナ名
	if env.Host == "" {
		return "", errors.New("HOST is not found")
	}
	env.User = os.Getenv("POSTGRES_USER")
	if env.User == "" {
		return "", errors.New("POSTGRES_USER is not found")
	}
	env.Pass = os.Getenv("POSTGRES_PASSWORD")
	if env.Pass == "" {
		return "", errors.New("POSTGRES_PASSWORD is not found")
	}
	env.DB = os.Getenv("POSTGRES_DB")
	if env.DB == "" {
		return "", errors.New("POSTGRES_DB is not found")
	}
	env.Port = os.Getenv("PORT")
	if env.Port == "" {
		env.Port = "5432"
	}
	env.TimeZone = os.Getenv("TIME_ZONE")
	if env.TimeZone == "" {
		env.TimeZone = "Asia/Tokyo"
	}
	
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=%s", env.Host, env.User, env.Pass, env.DB, env.Port, env.TimeZone)
	return dsn, nil
}

func Connect(dsn string) (*gorm.DB, error) {
	if dsn == "" {
		return &gorm.DB{}, errors.New("Dsn is not found, Please exec init")
	}
	var err error
	if db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
		if count < retry {
			count += 1

			log.Printf("Connect retry: %v", count)
			time.Sleep(time.Duration(count) * time.Second)
			db, err = Connect(dsn)
			if err != nil {
				return &gorm.DB{}, err
			}
			return db, nil
		}
		return &gorm.DB{}, errors.New("Failed DB connection: " + err.Error())
	}
	count = 0 // 接続回数リセット
	log.Printf("-- DB connected --")

	return db, nil
}
