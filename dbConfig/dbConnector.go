package dbConfig

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/lib/pq" // PostgreSQL driver
    "github.com/joho/godotenv"
)

func ConnectDB() *sql.DB {
    // Load environment variables
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    // Database connection string
    connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
    )

    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatalf("Error connecting to the database: %v", err)
    }

    // Check connection
    if err := db.Ping(); err != nil {
        log.Fatalf("Database connection failed: %v", err)
    }

    log.Println("Supabase PostgreSQL database connected successfully")
    return db
}




// package dbConfig

// import (
// 	"context"
// 	"fmt"
// 	"database/sql"
// 	"log"
// 	_ "github.com/lib/pq" // PostgreSQL driver
// )

// func ConnectToPostgreSQL() {
// 	fmt.Println("Connecting to Supabase PostgreSQL database")

// 	// Create a context
// 	ctx := context.Background()

// 	// Connect to MongoDB Atlas
// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI(DATABASE_URL))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	//Ping database
//     err = client.Ping(ctx, nil)
//     if err != nil {
//         log.Fatal(err)
//     }

//     fmt.Println("Successfully connected to mongo cluster")
// 	DATABASE = client.Database(DATABASE_NAME)
// }
