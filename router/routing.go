package router

import (

	// user defined packages
	"article/handler"
	"article/logs"
	"article/repository"

	// third party packages
	"github.com/gofiber/fiber/v2"
)

func Routing(db *repository.DbConnection) {
	h := handler.Newhandler(db)
	log := logs.Log()

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
	log.Info.Println("Message : 'Server starts in port 8000...' Status : 200")
	if err := app.Listen(":8000"); err != nil {
		log.Info.Println("Message : 'Error at start a server...' Status : 500")
		return
	}
}