package	 config

import (  
	"database/sql"
	"log" 
    "time"
	_ "github.com/go-sql-driver/mysql" 
    "app/internal/interfaces"
    "app/internal/common" 
    "fmt"
)
  
type DatabaseConfig struct {
	Driver   string
	Host     string
	Port     int
	DBName   string
	User     string
	Password string
	SSLMode  string
    GrpcAddr string // gRPC address for remote repository
}

// initMysql initializes the MySQL database connection
// It reads the database connection parameters from environment variables.
// It connects to the MySQL database and checks if the connection is successful.
func InitDB() (interfaces.DBStore, DatabaseConfig) {
    config := DatabaseConfig{
        Driver:   common.GetEnvOrDefault("DB_DRIVER", "mysql"),
        Host:     common.GetEnvOrDefault("DB_HOST", "localhost"),
        Port:     common.GetEnvIntOrDefault("DB_PORT", 3306),
        DBName:     common.GetEnvOrDefault("DB_DATABASE", "task_api"),
        User: common.GetEnvOrDefault("DB_USERNAME", "root"),
        Password: common.GetEnvOrDefault("DB_PASSWORD", ""),
        SSLMode:  common.GetEnvOrDefault("DB_SSL_MODE", "disable"),
        GrpcAddr: common.GetEnvOrDefault("DB_GRPC_ADDR", ":7070"),
    }
    // Db ,err := stores.NewDatabase(Database)

var dsn string
    
    switch config.Driver {
    case "postgres":
        dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
            config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)
    case "mysql":
        dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
            config.User, config.Password, config.Host, config.Port, config.DBName)
    case "local":
        return nil, config // Local repository does not require a database connection
    case "grpc":
        // gRPC does not require a database connection, return nil
        return nil, config
    // Add other database drivers as needed
    default:
         log.Fatal("unsupported database driver: %s", config.Driver)
    }

 fmt.Printf("Connecting to %s database at %s:%d...\n", config.Driver, config.Host, config.Port)
 fmt.Printf("DSN: %s\n", dsn)       
    // Open a new database connectio
    
    db, err := sql.Open(config.Driver, dsn)
    if err != nil {
        log.Fatal("failed to open database: %w", err) 
    }
    
   
    // var err error
    // Db, err := sql.Open("mysql", os.Getenv("MYSQL_ROOT_USER") + ":"+os.Getenv("MYSQL_ROOT_PASSWORD")+"@tcp("+os.Getenv("MYSQL_HOST") + ":3306"+")/"+os.Getenv("MYSQL_DATABASE")+"?parseTime=true")
   
    //  if err != nil {
    //     log.Fatal("Error connecting to DB:", err)
    // }

    // Configure connection pool
	db.SetMaxOpenConns(common.GetEnvIntOrDefault("DB_MAXOPEN",25)) // Max open connections
	db.SetMaxIdleConns(common.GetEnvIntOrDefault("DB_MAXIDLE",10))  // Max idle connections
	db.SetConnMaxLifetime(30 * time.Minute) // Max lifetime for a connection
	db.SetConnMaxIdleTime(10 * time.Minute) // Max idle time before closing
    

    if err = db.Ping(); err != nil {
        log.Fatal("Error pinging DB:", err)
    }
    var  dbStore interfaces.DBStore
    dbStore = db
    return dbStore, config
}

 