package services

import (
	"fmt"
	"github.com/revel/revel"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"test/app/models"
)

var Db *gorm.DB

func InitDB() {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel:                  logger.Error,
			IgnoreRecordNotFoundError: true,
		},
	)

	host := revel.Config.StringDefault("db.host", "localhost")
	dbname := revel.Config.StringDefault("db.name", "travel")
	user := revel.Config.StringDefault("db.user", "travel")
	password := revel.Config.StringDefault("db.password", "travel")

	connstring := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable", host, user, password, dbname)

	db, err := gorm.Open(postgres.Open(connstring), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "viewport.", // schema name
			SingularTable: false,
		},
	})

	Db = db

	if err != nil {
		// log.Errorf("Failed to load :%s",err.Error())
		logrus.Error(Db.Error)
	}

	logrus.Info("Database connected")

	//Create the schema for the migrations if in dev mode.  Production should
	//always already have this with the correct permissions
	if revel.RunMode == "dev" {
		logrus.Info("Creating viewport schema")
		err := Db.Exec("CREATE SCHEMA IF NOT EXISTS viewport").Error

		if err != nil {
			logrus.Error(err)
		}
	}

	logrus.Info("Running database migrations")
	err = Db.AutoMigrate(
		&models.WizProject{},
		&models.WizSubscription{},
		&models.AwsAccount{},
		&models.Application{},
		&models.Team{},
		&models.User{},
		&models.UserTeamMap{},
		&models.WizApplication{},
		&models.GhaHistory{},
	)

	if err != nil {
		logrus.Error(err)
	}

}
