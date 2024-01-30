package router

import (

	// user defined packages
	"article/handler"
	"article/helpers"
	"article/repository"
	"time"

	// third party packages
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func Routing(db *repository.DbConnection) {
	h := handler.Newhandler(db)
	helpers.Log.WithFields(logrus.Fields{
		"service": "article", "function": "routing", "started_at": time.Now(),
	}).Info("Routing started")

	//Initializing the fiber router
	app := fiber.New()

	//Routes for comments
	commentRoutes := app.Group("/article-exe/v1/comment")
	commentRoutes.Post("/add-comment/", h.AddComment)
	commentRoutes.Get("/list-all-comments/", h.LisAllComments)
	commentRoutes.Post("/add-comment-on-comment/", h.AddCommentOnComment)

	//Routes for articles
	articleRoutes := app.Group("/article-exe/v1/article")
	articleRoutes.Post("/add-article/", h.AddArticle)
	articleRoutes.Get("/list-all-articles", h.ListAllArticles)
	articleRoutes.Get("/list-article-by-id/", h.ListArticleByID)

	//Starting the server
	helpers.Log.WithFields(logrus.Fields{
		"service": "article",
		"message": "Port started at 8000!",
	}).Info("Message : 'Server starts in port 8000...' Status : 200")
	if err := app.Listen(":8000"); err != nil {
		helpers.Log.WithFields(logrus.Fields{
			"service": "article",
			"error":   err.Error(),
		}).Error("Error at start a server... Status : 500")
		return
	}
}
