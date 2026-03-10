package main

import (
	"ta-karangtaruna/database"
	controllers "ta-karangtaruna/internal/controller"
	"ta-karangtaruna/internal/entities"
	"ta-karangtaruna/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {

	// CONNECT DATABASE
	database.ConnectDatabase()

	// AUTO MIGRATE
	database.DB.AutoMigrate(
		&entities.User{},
		&entities.Kategori{},
		&entities.Kegiatan{},
		&entities.Komentar{},
		&entities.Dokumentasi{},
	)

	r := gin.Default()
	r.Static("/uploads", "./uploads")

	// ========================
	// PUBLIC ROUTES
	// ========================

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API Karang Taruna berjalan 🚀",
		})
	})

	r.POST("/register", controllers.Register)
	r.POST("/register-ketua", controllers.RegisterKetua)
	r.POST("/login", controllers.Login)

	// publik bisa lihat kegiatan
	r.GET("/kegiatan", controllers.GetAllKegiatan)
	r.GET("/kegiatan/:id/komentar", controllers.GetKomentar)

	r.GET("/kegiatan/:id/dokumentasi", controllers.GetDokumentasi)

	// ========================
	// USER ROUTES (LOGIN REQUIRED)
	// ========================

	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		api.GET("/profile", controllers.GetProfile)

		// anggota membuat kegiatan
		api.POST("/kegiatan", controllers.CreateKegiatan)

		// kegiatan milik user login
		api.GET("/kegiatan-saya", controllers.GetKegiatanSaya)
		api.POST("/kegiatan/:id/komentar", controllers.CreateKomentar)

		api.POST("/kegiatan/:id/dokumentasi", controllers.UploadDokumentasi)
	}

	// ========================
	// KETUA ROUTES
	// ========================

	ketua := r.Group("/api/ketua")
	ketua.Use(middleware.AuthMiddleware(), middleware.OnlyKetua())
	{
		ketua.GET("/dashboard", controllers.GetDashboardKetua)

		// melihat semua user
		ketua.GET("/users", controllers.GetAllUsers)

		// melihat semua kegiatan
		ketua.GET("/kegiatan", controllers.GetAllKegiatanKetua)
		ketua.GET("/kegiatan/user/:id", controllers.GetKegiatanByUser)
		ketua.PATCH("/kegiatan/:id/approve", controllers.ApproveKegiatan)
		ketua.PATCH("/kegiatan/:id/reject", controllers.RejectKegiatan)
	}

	r.Run(":8080")
}
