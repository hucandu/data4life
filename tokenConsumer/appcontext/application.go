package appcontext

import (
	"fmt"
	"github.com/hucandu/data4life/tokenConsumer/models"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
)

type AppContext struct {
	DbClient *gorm.DB
}

type conf struct {
	DbPort     string `yaml:db_host`
	DbName     string `yaml:db_port`
	DbUser     string `yaml:db_name`
	DbHost     string `yaml:db_user`
	DbPassword string `yaml:db_password`
}

func getConf(env string) *conf {
	yamlFile, err := ioutil.ReadFile("conf/" + env + ".yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	var c *conf
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}

func Initiate(env string) *AppContext {
	appContext := &AppContext{
		DbClient: setupDatabase(env),
	}
	return appContext
}

func setupDatabase(env string) *gorm.DB {
	dbConf := getConf(env)
	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		dbConf.DbUser, dbConf.DbPassword, dbConf.DbHost, dbConf.DbPort, dbConf.DbName)

	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&models.TokenData{})
	return db
}
