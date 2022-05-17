package appcontext

import (
	"fmt"
	"github.com/hucandu/data4life/tokenConsumer/models"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"path/filepath"
)

type AppContext struct {
	DbClient *gorm.DB
}

type conf struct {
	DbConf struct {
		DbPort     string `yaml:"db_port"`
		DbName     string `yaml:"db_name"`
		DbUser     string `yaml:"db_user"`
		DbHost     string `yaml:"db_host"`
		DbPassword string `yaml:"db_password"`
	} `yaml:"dbConf"`
}

func getConf(env string) *conf {
	filePath, err := filepath.Abs("conf/" + env + ".yaml")
	if err != nil {
		log.Fatalf(err.Error())
		return nil
	}
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	c := &conf{}
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
	conf := getConf(env)
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		conf.DbConf.DbUser, conf.DbConf.DbPassword, conf.DbConf.DbHost, conf.DbConf.DbPort, conf.DbConf.DbName)

	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&models.TokenData{})
	db.Exec(fmt.Sprintf("TRUNCATE TABLE %s;", "token_data"))
	return db
}
