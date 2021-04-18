package main

import (
	"log"
	"net/http"
	"rental-buku/auth"
	"rental-buku/book"
	"rental-buku/category"
	"rental-buku/handler"
	"rental-buku/helper"
	"rental-buku/seeds"
	"rental-buku/user"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	db := helper.SetupDB()

	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	for _, seed := range seeds.All() {
		if err := seed.Run(db); err != nil {
			log.Printf("Running seed '%s', failed with error: %s", seed.Name, err)
		}
	}

	// Repository
	userRepository := user.NewRepository(db)
	categoryRepository := category.NewRepository(db)
	bookRepository := book.NewRepository(db)

	// Service
	authService := auth.NewService()
	userService := user.NewService(userRepository)
	categoryService := category.NewService(categoryRepository)
	bookService := book.NewService(bookRepository)

	// Handler
	userHandler := handler.NewUserHandler(userService, authService)
	categoryHandler := handler.NewCategoryHandler(categoryService, authService)
	bookHandler := handler.NewBookHandler(bookService, authService)

	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/photos", authMiddleware(authService, userService), userHandler.UploadPhoto)

	api.GET("/categories", categoryHandler.GetCategories)
	api.GET("/categories/:id", categoryHandler.GetCategoryById)
	api.POST("/categories", authMiddleware(authService, userService), categoryHandler.CreateCategory)
	api.PUT("/categories/:id", authMiddleware(authService, userService), categoryHandler.UpdateCategory)
	api.DELETE("/categories/:id", authMiddleware(authService, userService), categoryHandler.DeleteCategory)

	api.GET("/books", bookHandler.GetBooks)
	api.GET("/books/:id", bookHandler.GetBookById)
	api.POST("/books", authMiddleware(authService, userService), bookHandler.CreateBook)
	api.PUT("/books/:id", authMiddleware(authService, userService), bookHandler.UpdateBook)
	api.DELETE("/books/:id", authMiddleware(authService, userService), bookHandler.DeleteBook)

	router.Use(cors.Default())

	router.Run()
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)

		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userId := int(claim["user_id"].(float64))

		user, err := userService.GetUserById(userId)

		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}
