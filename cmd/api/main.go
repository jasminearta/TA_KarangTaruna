// @title API Sistem Informasi Karang Taruna
// @version 1.0
// @description Dokumentasi API backend Sistem Informasi Karang Taruna
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
package main

import (
	"ta-karangtaruna/database"
	docs "ta-karangtaruna/docs"
	controllers "ta-karangtaruna/internal/controller"
	"ta-karangtaruna/internal/entities"
	"ta-karangtaruna/internal/middleware"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDatabase()

	database.DB.AutoMigrate(
		&entities.User{},
		&entities.Kategori{},
		&entities.Inovasi{},
		&entities.Kegiatan{},
		&entities.FotoKegiatan{},
		&entities.Notification{},
		&entities.ApprovalLog{},
		&entities.FotoInovasi{},
	)

	r := gin.Default()
	r.Static("/uploads", "./uploads")

	docs.SwaggerInfo.Title = "API Sistem Informasi Karang Taruna"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/"

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API Karang Taruna berjalan 🚀",
		})
	})

	r.POST("/register", controllers.Register)
	r.POST("/register-ketua", controllers.RegisterKetua)
	r.POST("/login", controllers.Login)

	// ========================
	// PUBLIC ROUTES
	// ========================

	r.GET("/kegiatan", controllers.GetAllKegiatan)
	r.GET("/kegiatan/:id", controllers.GetDetailKegiatan)
	r.GET("/kegiatan/:id/foto", controllers.GetFotoKegiatan)

	r.GET("/inovasi", controllers.GetAllInovasi)
	r.GET("/inovasi/:id", controllers.GetDetailInovasi)
	r.GET("/inovasi/:id/foto", controllers.GetFotoInovasi)

	// ========================
	// AUTH ROUTES (LOGIN REQUIRED)
	// ========================

	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		api.GET("/profile", controllers.GetProfile)
		api.PUT("/profile", controllers.UpdateProfile)
		api.PUT("/profile/password", controllers.ChangePassword)
		api.POST("/profile/foto", controllers.UploadFotoProfile)

		api.GET("/notifications", controllers.GetNotifications)
		api.PATCH("/notifications/:id/read", controllers.ReadNotification)
	}

	// ========================
	// KETUA DIVISI ROUTES
	// ========================

	ketuaDivisi := r.Group("/api")
	ketuaDivisi.Use(middleware.AuthMiddleware(), middleware.OnlyKetuaDivisi())
	{
		ketuaDivisi.POST("/kegiatan", controllers.CreateKegiatan)
		ketuaDivisi.GET("/kegiatan-saya", controllers.GetKegiatanSaya)
		ketuaDivisi.PUT("/kegiatan/:id", controllers.UpdateKegiatan)
		ketuaDivisi.POST("/kegiatan/:id/foto", controllers.UploadFotoKegiatan)
		ketuaDivisi.DELETE("/kegiatan/:id", controllers.DeleteKegiatan)

		ketuaDivisi.POST("/inovasi", controllers.CreateInovasi)
		ketuaDivisi.GET("/inovasi-saya", controllers.GetInovasiSaya)
		ketuaDivisi.PUT("/inovasi/:id", controllers.UpdateInovasi)
		ketuaDivisi.DELETE("/inovasi/:id", controllers.DeleteInovasi)
		ketuaDivisi.POST("/inovasi/:id/foto", controllers.UploadFotoInovasi)
	}

	// ========================
	// KETUA UMUM ROUTES
	// ========================

	ketuaUmum := r.Group("/api/ketua")
	ketuaUmum.Use(middleware.AuthMiddleware(), middleware.OnlyKetuaUmum())
	{
		ketuaUmum.GET("/dashboard", controllers.GetDashboardKetua)

		ketuaUmum.GET("/users", controllers.GetAllUsers)

		ketuaUmum.GET("/kegiatan", controllers.GetAllKegiatanKetua)
		ketuaUmum.GET("/kegiatan/user/:id", controllers.GetKegiatanByUser)
		ketuaUmum.PATCH("/kegiatan/:id/approve", controllers.ApproveKegiatan)
		ketuaUmum.PATCH("/kegiatan/:id/reject", controllers.RejectKegiatan)

		ketuaUmum.GET("/inovasi", controllers.GetAllInovasiKetua)
		ketuaUmum.GET("/inovasi/user/:id", controllers.GetInovasiByUser)
		ketuaUmum.PATCH("/inovasi/:id/approve", controllers.ApproveInovasi)
		ketuaUmum.PATCH("/inovasi/:id/reject", controllers.RejectInovasi)
	}

	r.Run(":8080")
}
