package router

import (
	"go-fiber-boilerplate/internal/config"
	"go-fiber-boilerplate/internal/handlers"
	"go-fiber-boilerplate/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Setup(app *fiber.App, db *gorm.DB, cfg *config.Config) {
	// Initialize handlers
	staticHandler := handlers.NewStaticHandler()
	authHandler := handlers.NewAuthHandler(db, cfg)
	userHandler := handlers.NewUserHandler(db)
	micropostHandler := handlers.NewMicropostHandler(db)
	relationshipHandler := handlers.NewRelationshipHandler(db)

	// Static pages
	app.Get("/", staticHandler.Home)
	app.Get("/about", staticHandler.About)
	app.Get("/help", staticHandler.Help)
	app.Get("/contact", staticHandler.Contact)

	// Authentication routes
	app.Get("/signup", authHandler.SignupForm)
	app.Post("/signup", authHandler.Signup)
	app.Get("/login", authHandler.LoginForm)
	app.Post("/login", authHandler.Login)
	app.Post("/logout", authHandler.Logout)

	// Account activation
	app.Get("/account_activations/:id/edit", authHandler.ActivateAccount)

	// Password reset
	app.Get("/password_resets/new", authHandler.PasswordResetForm)
	app.Post("/password_resets", authHandler.CreatePasswordReset)
	app.Get("/password_resets/:id/edit", authHandler.EditPasswordReset)
	app.Patch("/password_resets/:id", authHandler.UpdatePasswordReset)

	// User routes
	app.Get("/users", userHandler.Index)
	app.Get("/users/:id", userHandler.Show)

	// Protected user routes
	protected := app.Group("", middleware.RequireAuth(db, cfg.JWTSecret))
	protected.Get("/users/:id/edit", userHandler.EditForm)
	protected.Patch("/users/:id", userHandler.Update)
	protected.Delete("/users/:id", userHandler.Delete)
	protected.Get("/users/:id/following", userHandler.Following)
	protected.Get("/users/:id/followers", userHandler.Followers)

	// Micropost routes
	protected.Post("/microposts", micropostHandler.Create)
	protected.Delete("/microposts/:id", micropostHandler.Delete)

	// Relationship routes
	protected.Post("/relationships", relationshipHandler.Create)
	protected.Delete("/relationships/:id", relationshipHandler.Delete)

	// API routes
	api := app.Group("/api/v1")

	// API Authentication
	api.Post("/signup", authHandler.APISignup)
	api.Post("/login", authHandler.APILogin)
	api.Post("/logout", middleware.RequireAuth(db, cfg.JWTSecret), authHandler.APILogout)

	// API Users
	apiProtected := api.Group("", middleware.RequireAuth(db, cfg.JWTSecret))
	api.Get("/users", userHandler.APIIndex)
	api.Get("/users/:id", userHandler.APIShow)
	apiProtected.Patch("/users/:id", userHandler.APIUpdate)
	apiProtected.Delete("/users/:id", userHandler.APIDelete)

	// API Microposts
	apiProtected.Get("/microposts", micropostHandler.APIIndex)
	apiProtected.Post("/microposts", micropostHandler.APICreate)
	apiProtected.Get("/microposts/:id", micropostHandler.APIShow)
	apiProtected.Patch("/microposts/:id", micropostHandler.APIUpdate)
	apiProtected.Delete("/microposts/:id", micropostHandler.APIDelete)

	// API Relationships
	apiProtected.Get("/relationships", relationshipHandler.APIIndex)
	apiProtected.Post("/relationships", relationshipHandler.APICreate)
	apiProtected.Delete("/relationships/:id", relationshipHandler.APIDelete)
}
