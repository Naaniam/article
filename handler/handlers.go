package handler

import (

	//user defined packages

	"article/logs"
	"article/models"
	"article/repository"
	"article/utilities"
	"fmt"

	"strconv"

	//built-in packages

	"time"

	//third-party package

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Repo *repository.DbConnection
}

func Newhandler(db *repository.DbConnection) *Handler {
	return &Handler{Repo: db}
}

//-------------------------------Article-------------------------------------------------------------------------------

// Handler function to add articles
func (h *Handler) AddArticle(c *fiber.Ctx) error {
	article := models.Article{}
	log := logs.Log()

	log.Info.Println("Message : 'AddArticle-API called'")
	if err := c.BodyParser(&article); err != nil {
		log.Error.Printf("Error : %s, Status : 400", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	article.CreationDate = time.Now()

	if err := h.Repo.AddArticle(&article); err != nil {
		log.Error.Printf("Error :%s  Status : 400", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Publish message to Kafka
	go utilities.PublishToKafka(fmt.Sprintf("%v", article), "articles.create", article.ID)

	log.Info.Println("Message : 'Article added successfully' Status : 201")
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Created Article!", "ArticleID": article.ID})
}

// Handler function to listall the aticles in golang
func (h *Handler) ListAllArticles(c *fiber.Ctx) error {
	articles := []models.Article{}
	log := logs.Log()

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < 1 {
		page = 1
	}

	perPage := 20
	offset := (page - 1) * perPage

	log.Info.Println("Message : 'ListAllArticles-API called'")
	if err := h.Repo.ListAllArticles(perPage, offset, &articles); err != nil {
		log.Error.Printf("Error : %s Status : 404", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Publish message to Kafka
	for _, article := range articles {
		go utilities.PublishToKafka(fmt.Sprintf("%v", article), "articles.list-all-articlesss", article.ID)
	}

	log.Info.Println("Message : 'Article(s) retrieved successfully' Status : 200")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "retrived successfully", "articles": articles})
}

// Handler function to list the article with ID
func (h *Handler) ListArticleByID(c *fiber.Ctx) error {
	article := models.Article{}
	log := logs.Log()

	log.Info.Println("Message : 'ListArticleByID-API called'")
	if err := h.Repo.ListArticleByID(c.Query("article_id"), &article); err != nil {
		log.Error.Printf("Error : %s Status : 400", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	go utilities.PublishToKafka(fmt.Sprintf("%v", article), "articles.list-article-by-id", article.ID)

	log.Info.Println("Message : 'Article retrieved successfully' Status : 200")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"article": article})
}

// -------------------------------------Comments------------------------------------------------------------------------

// Handler function to add comments
func (h *Handler) AddComment(c *fiber.Ctx) error {
	comment := models.Comment{}
	log := logs.Log()

	log.Info.Println("Message : 'AddComment-API called'")

	//parsing comment data from request body
	if err := c.BodyParser(&comment); err != nil {
		log.Error.Printf("Error : %s Status : 400", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	comment.CreationDate = time.Now()
	if err := h.Repo.AddComment(c.Query("article_id"), &comment); err != nil {
		log.Error.Printf("Error : %s Status : 400", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	go utilities.PublishToKafka(fmt.Sprintf("%v", comment), "comment.create", comment.ID)

	log.Info.Println("Message : 'Added comment successfully' Status : 201")
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"Message": "Added comment successfully", "CommentID": comment.ID})
}

// Handler function to list all comments
func (h *Handler) LisAllComments(c *fiber.Ctx) error {
	comments := []models.Comment{}
	log := logs.Log()

	log.Info.Println("Message : 'LisAllComments-API called'")

	if err := h.Repo.ListAllComments(c.Query("article_id"), &comments); err != nil {
		log.Error.Printf("Error : %s Status : 400", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	for _, comment := range comments {
		go utilities.PublishToKafka(fmt.Sprintf("%v", comment), "articles.get-comments-by-post-id", comment.ID)
	}

	log.Info.Println("Message : 'Comment(s) retrieved successfully' Status : 200")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "retrived all the comments", "comments": comments})
}

// Handler fucntion to comment on comment
func (h *Handler) AddCommentOnComment(c *fiber.Ctx) error {
	reply := models.Reply{}
	log := logs.Log()

	log.Info.Println("Message : 'AddCommentOnComment-API called'")

	if err := c.BodyParser(&reply); err != nil {
		log.Error.Printf("Error : %s Status : 400", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	reply.CreationDate = time.Now()

	if err := h.Repo.AddCommentOnComment(c.Query("comment_id"), &reply); err != nil {
		log.Error.Printf("Error : %s Status : 400", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	go utilities.PublishToKafka(fmt.Sprintf("%v", reply), "articles.post-comment-on-comment", reply.ID)

	log.Info.Println("Message : 'Added comment(s) on comment' Status : 200")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "Added comment on comment successfully"})
}
