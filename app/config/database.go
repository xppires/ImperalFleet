package	 config

import (  
	"database/sql"
	"log"
	"os" 
    "time"
	_ "github.com/go-sql-driver/mysql" 
    "app/internal/interfaces"
    "app/internal/common" 
)
  
type DatabaseConfig struct {
	Driver   string
	Host     string
	Port     int
	Name     string
	Username string
	Password string
	SSLMode  string
}

// initMysql initializes the MySQL database connection
// It reads the database connection parameters from environment variables.
// It connects to the MySQL database and checks if the connection is successful.
func InitMysql() interfaces.DBStore {
    // Database := DatabaseConfig{
    //     Driver:   common.GetEnvOrDefault("DB_DRIVER", "mysql"),
    //     Host:     common.GetEnvOrDefault("DB_HOST", "localhost"),
    //     Port:     common.GetEnvIntOrDefault("DB_PORT", 3306),
    //     Name:     common.GetEnvOrDefault("DB_NAME", "task_api"),
    //     Username: common.GetEnvOrDefault("DB_USERNAME", "root"),
    //     Password: common.GetEnvOrDefault("DB_PASSWORD", ""),
    //     SSLMode:  common.GetEnvOrDefault("DB_SSL_MODE", "disable"),
    // }
    // Db ,err := stores.NewDatabase(Database)

    // var err error
    Db, err := sql.Open("mysql", os.Getenv("MYSQL_ROOT_USER") + ":"+os.Getenv("MYSQL_ROOT_PASSWORD")+"@tcp("+os.Getenv("MYSQL_HOST") + ":3306"+")/"+os.Getenv("MYSQL_DATABASE")+"?parseTime=true")
   
     if err != nil {
        log.Fatal("Error connecting to DB:", err)
    }

    // Configure connection pool
	Db.SetMaxOpenConns(common.GetEnvIntOrDefault("DB_MAXOPEN",25)) // Max open connections
	Db.SetMaxIdleConns(common.GetEnvIntOrDefault("DB_MAXIDLE",10))  // Max idle connections
	    Db.SetConnMaxLifetime(30 * time.Minute) // Max lifetime for a connection
	Db.SetConnMaxIdleTime(10 * time.Minute) // Max idle time before closing
    

    if err = Db.Ping(); err != nil {
        log.Fatal("Error pinging DB:", err)
    }
    var  dbStore interfaces.DBStore
    dbStore = Db
    return dbStore
}

 