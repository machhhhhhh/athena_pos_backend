package configs

// import (
// 	"os"
// 	"time"

// 	"gorm.io/driver/sqlserver"
// 	"gorm.io/gorm"
// )

// var DB *gorm.DB

// func Connection() {
// 	host := os.Getenv("E_MEMO_DB_HOST")
// 	user := os.Getenv("E_MEMO_DB_USERNAME")
// 	password := os.Getenv("E_MEMO_DB_PASSWORD")
// 	dbname := os.Getenv("E_MEMO_DB_NAME")
// 	dbport := os.Getenv("E_MEMO_DB_PORT")

// 	// Check if the environment variable is set
// 	if host == "" || user == "" || password == "" || dbname == "" || dbport == "" {
// 		panic("Environment variable is not set")
// 	}

// 	dsn := "sqlserver://" + user + ":" + password + "@" + host + ":" + dbport + "?database=" + dbname
// 	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		panic("Failed to connect to the database: " + err.Error())
// 	}

// 	sqlDb, err := db.DB()
// 	if err != nil {
// 		panic(err)
// 	}

// 	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
// 	sqlDb.SetMaxIdleConns(10)

// 	// SetMaxOpenConns sets the maximum number of open connections to the database.
// 	sqlDb.SetMaxOpenConns(151)

// 	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
// 	// Reccommented it's should lower than db config and by servive time.
// 	sqlDb.SetConnMaxLifetime(time.Minute * 10)

// 	DB = db
// }
