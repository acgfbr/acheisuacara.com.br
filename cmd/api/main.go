package main

import (
	"fmt"
	"log"

	"acheisuacara.com.br/pkg/config"
	"acheisuacara.com.br/pkg/database"
	"acheisuacara.com.br/pkg/handlers"
	"acheisuacara.com.br/pkg/middleware"
	"acheisuacara.com.br/pkg/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize database
	db, err := database.NewMySQLConnection(&cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize Redis client
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	// Initialize services and handlers
	urlService := services.NewURLService(db)
	urlHandler := handlers.NewURLHandler(urlService)
	rateLimiter := middleware.NewRateLimiter(redisClient, cfg.Server.RateLimit, cfg.Server.RateInterval)

	// Setup router
	router := gin.Default()

	// Configure CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Apply rate limiter middleware
	router.Use(rateLimiter.Middleware())

	// Routes
	router.POST("/api/shorten", urlHandler.CreateShortURL)
	router.GET("/:shortCode", urlHandler.RedirectToLongURL)

	// Start server
	log.Printf("Server starting on port %s", cfg.Server.Port)
	if err := router.Run(cfg.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
