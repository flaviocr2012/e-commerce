    package main

    import (
        "database/sql"
        "e-commerce/controller"
        "e-commerce/repository"
        "e-commerce/router"
        "e-commerce/service"
        "log"
        "net/http"

        "github.com/golang-migrate/migrate/v4"
        _ "github.com/golang-migrate/migrate/v4/database/postgres"
        _ "github.com/golang-migrate/migrate/v4/source/file"
    )

    func main() {
        log.Println("Starting application")

        // Database connection
        connStr := "user=flavio password=securepassword dbname=flaviodb host=localhost port=5433 sslmode=disable"

        db, err := sql.Open("postgres", connStr)
        if err != nil {
            log.Fatal("DB Connection failed:", err)
        }
        defer db.Close()

        log.Println("Database connected successfully")

        // Initialize your repository
        userRepository := repository.NewUserRepository(db)
        productRepository := repository.NewProductRepository(db)

        // Initialize your services
        userService := service.NewUserService(userRepository)
        productService := service.NewProductService(productRepository)

        // Initialize your controllers
        userController := controller.NewUserController(userService)
        productController := controller.NewProductController(productService)

        // Set up the router
        r := router.SetUpRouter(userController, productController)

        // Start the server
        log.Println("Server starting on: 8080")
        if err := http.ListenAndServe(":8080", r); err != nil {
            log.Fatal("Server Failed", err)
        }
    }

    // runMigrations handles database migrations
    func runMigrations() error {
        m, err := migrate.New(
            "file://db/migrations",
            "postgres://flavio:securepassword@localhost:5433/flaviodb?sslmode=disable")
        if err != nil {
            return err
        }

        if err := m.Up(); err != nil && err != migrate.ErrNoChange {
            return err
        }

        return nil
    }