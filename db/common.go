package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/dosorio55/gambitApi/models"
	"github.com/dosorio55/gambitApi/secretManager"
  _ "github.com/go-sql-driver/mysql"
)

var SecretModel models.SecretStructJSON
var err error
var Db *sql.DB

func ReadSecret() error {
	//this os.Getenv("SecretName") is the name of the secret in AWS Secrets Manager
	// that is probably going to be in the .env file
	// and it is going to be de name of the database
	SecretModel, err = secretmanager.GetSecret(os.Getenv("SecretName"))
	return err
}

func DbConnect() error {
	Db, err = sql.Open("mysql", ConnStr(SecretModel))

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = Db.Ping()

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Successfully connected to database!")
	return nil
}

func ConnStr(data models.SecretStructJSON) string {
	var dbUser, authToken, dbEndPoint, dbName string
	dbUser = data.Username
	authToken = data.Password
	dbEndPoint = data.Host
	dbName = "gambit"
	// dbName = data.DbClusterIdentifier

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowClearPasswords=true", dbUser, authToken, dbEndPoint, dbName)
	fmt.Println("dsn: ", dsn)

	return dsn
}
