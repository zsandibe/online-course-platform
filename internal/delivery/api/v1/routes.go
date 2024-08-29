package v1

import "github.com/gin-gonic/gin"

func (h *Handler) Routes() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Recovery())

	api := router.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", h.signUp)
			auth.POST("/register", h.signIn)
			auth.POST("/forgot-password", h.forgotPassword)
			auth.POST("/verify-email", h.verifyEmail)
		}
		profile := api.Group("/profile/:nickname")
		{
			profile.GET("") // TODO: user stats
			settings := profile.Group("/settings")
			{
				settings.PUT(":id", h.updateProfile)
				settings.DELETE(":id", h.deleteAccount)
				settings.GET("/change-password", h.changePassword)
			}

		}
		course := api.Group("/course")
		{
			course.POST("")       // TODO: create a new course || ONLY TEACHER
			course.GET("/:id")    // TODO: get a course by ID
			course.PUT("/:id")    // TODO: update a course by ID || ONLY TEACHER
			course.DELETE("/:id") // TODO: delete a course by ID || ONLY TEACHER
			course.POST("/:id")   // TODO: enroll to course || ONLY STUDENT
			course.DELETE("/:id") // TODO: leave from course || ONLY STUDENT
		}

		lesson := api.Group("/lesson/:course_id")
		{
			lesson.POST("")       // TODO: create a new lesson || ONLY TEACHER
			lesson.GET("/:id")    // TODO: get a lesson by ID
			lesson.PUT("/:id")    // TODO: update a lesson by ID || ONLY TEACHER
			lesson.DELETE("/:id") // TODO: delete a lesson by ID || ONLY TEACHER
		}

		category := api.Group("/category")
		{
			category.GET("")        // TODO: get all categories
			category.POST("")       // TODO: create a new category || ONLY ADMIN
			category.GET("/:id")    // TODO: get a category by ID
			category.PUT("/:id")    // TODO: update a category by ID || ONLY ADMIN
			category.DELETE("/:id") // TODO: delete a category by ID || ONLY ADMIN
		}

		enrollment := api.Group("/enrollment")
		{
			enrollment.GET("") // TODO: get all enrollments history stats ||  ONLY ADMIN

			action := api.Group("/:course_id")
			{
				action.GET("")        // TODO: get all enrollments for a course
				action.GET("/:id")    // TODO: get an enrollment by ID
				action.DELETE("/:id") // TODO: cancel an enrollment by ID || ONLY STUDENT
			}

		}

		payment := api.Group("/payment/:enrollment_id")
		{
			payment.POST("")   // TODO: make a payment for an enrollment || ONLY STUDENT
			payment.GET(":id") // TODO: get a payment by ID

		}
	}
	return router
}
