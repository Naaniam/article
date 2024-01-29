package handler

import (
	"article/drivers"
	"article/repository"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

//------------------------------------------------Article-------------------------------------------------------------

// test for the handler to add article
func TestAddArticle(t *testing.T) {
	connection := drivers.TestSQLDriver()
	database := repository.DbConnection{
		DB: connection,
	}
	db := Newhandler(&database)

	app := fiber.New()

	// Define the route using the AddArticle handler
	app.Post("/article-exe/v1/article/add-article", db.AddArticle)

	// Define the test case
	t.Run("missing nickname", func(t *testing.T) {
		body := `{
			"nickname": "",
			"title": "Article-1",
			"content": "This is the content of the Article-1."
		}`
		req := httptest.NewRequest(http.MethodPost, "/article-exe/v1/article/add-article", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		statusCode, err := strconv.Atoi(resp.Status[:3])
		assert.NoError(t, err)

		// Assert the response status code
		assert.Equal(t, fiber.StatusBadRequest, statusCode)
	})

	t.Run("missing title", func(t *testing.T) {
		body := `{
			"nickname": "Haritha",
			"title": "",
			"content": "This is the content of the Article-1."
		}`
		req := httptest.NewRequest(http.MethodPost, "/article-exe/v1/article/add-article", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		statusCode, err := strconv.Atoi(resp.Status[:3])
		assert.NoError(t, err)

		// Assert the response status code
		assert.Equal(t, fiber.StatusBadRequest, statusCode)
	})

	t.Run("missing content", func(t *testing.T) {
		body := `{
			"nickname": "Haritha",
			"title": "Article-1",
			"content": ""
		}`
		req := httptest.NewRequest(http.MethodPost, "/article-exe/v1/article/add-article", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		statusCode, err := strconv.Atoi(resp.Status[:3])
		assert.NoError(t, err)

		// Assert the response status code
		assert.Equal(t, fiber.StatusBadRequest, statusCode)
	})

	t.Run("Product added successfully", func(t *testing.T) {
		body := `{
			"nickname": "Haritha",
			"title": "Article-1",
			"content": "This is the content of the Article-1."
		}`

		for i := 0; i < 2; i++ {
			req := httptest.NewRequest(http.MethodPost, "/article-exe/v1/article/add-article", strings.NewReader(body))
			req.Header.Set("Content-Type", "application/json")

			resp, err := app.Test(req)
			statusCode, err := strconv.Atoi(resp.Status[:3])
			assert.NoError(t, err)

			// Assert the response status code
			assert.Equal(t, fiber.StatusCreated, statusCode)
		}
	})
}

// test for the handler to list all the articles
func TestListAllArticles(t *testing.T) {
	connection := drivers.TestSQLDriver()
	db := Handler{
		Repo: &repository.DbConnection{DB: connection},
	}

	app := fiber.New()

	// Define the route using the AddArticle handler
	app.Get("/article-exe/v1/article/list-all-articles", db.ListAllArticles)

	// Define the test case

	t.Run("All Articles are retrieved successfully", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/article-exe/v1/article/list-all-articles", nil)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		statusCode, err := strconv.Atoi(resp.Status[:3])
		assert.NoError(t, err)

		// Assert the response status code
		assert.Equal(t, fiber.StatusOK, statusCode)
	})
}

// test for the handler to list the article based on ID
func TestListArticleByID(t *testing.T) {
	connection := drivers.TestSQLDriver()
	db := Handler{
		Repo: &repository.DbConnection{DB: connection},
	}

	app := fiber.New()

	// Define the route using the AddArticle handler
	app.Get("/article-exe/v1/article/list-article-by-id/", db.ListArticleByID)

	// Define the test case

	t.Run("Article retrieved successfully", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/article-exe/v1/article/list-article-by-id/?article_id=1", nil)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		statusCode, err := strconv.Atoi(resp.Status[:3])
		assert.NoError(t, err)

		// Assert the response status code
		assert.Equal(t, fiber.StatusOK, statusCode)
	})
}

//-----------------------------------------testing for comments------------------------------------------------------

// test for the handler to add comment
func TestAddComment(t *testing.T) {
	connection := drivers.TestSQLDriver()
	db := Handler{
		Repo: &repository.DbConnection{DB: connection},
	}

	app := fiber.New()

	// Define the route using the AddArticle handler
	app.Post("/article-exe/v1/comment/add-comment/", db.AddComment)

	// Define the test case
	t.Run("missing nickname", func(t *testing.T) {
		body := `{
			{
				"nickname": "",
				"content": "Great article!"
			}
		}`
		req := httptest.NewRequest(http.MethodPost, "/article-exe/v1/comment/add-comment/?article_id=1", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		statusCode, err := strconv.Atoi(resp.Status[:3])
		assert.NoError(t, err)

		// Assert the response status code
		assert.Equal(t, fiber.StatusBadRequest, statusCode)
	})

	t.Run("missing content", func(t *testing.T) {
		body := `{
			{
				"nickname": "Gnanitha",
				"content": ""
			}
		}`

		req := httptest.NewRequest(http.MethodPost, "/article-exe/v1/comment/add-comment/?article_id=1", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		statusCode, err := strconv.Atoi(resp.Status[:3])
		assert.NoError(t, err)

		// Assert the response status code
		assert.Equal(t, fiber.StatusBadRequest, statusCode)
	})

	t.Run("comment added successfully", func(t *testing.T) {
		body := `{
			{
				"nickname": "Gnanitha",
				"content": "Great article!!"
			}
		}`

		req := httptest.NewRequest(http.MethodPost, "/article-exe/v1/comment/add-comment/?article_id=1", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		statusCode, err := strconv.Atoi(resp.Status[:3])
		assert.NoError(t, err)

		// Assert the response status code
		assert.Equal(t, fiber.StatusBadRequest, statusCode)
	})
}

// test for handler to list all the comments for partiular post id
func TestLisAllComments(t *testing.T) {
	connection := drivers.TestSQLDriver()
	db := Handler{
		Repo: &repository.DbConnection{DB: connection},
	}

	app := fiber.New()

	// Define the route using the AddArticle handler
	app.Get("/article-exe/v1/comment/list-all-comments/", db.ListArticleByID)

	// Define the test case

	t.Run("Article retrieved successfully", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/article-exe/v1/comment/list-all-comments/?article_id=1", nil)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		statusCode, err := strconv.Atoi(resp.Status[:3])
		assert.NoError(t, err)

		// Assert the response status code
		assert.Equal(t, fiber.StatusOK, statusCode)
	})
}

// test for handler to post comment on comment
func TestAddCommentOnComment(t *testing.T) {
	connection := drivers.TestSQLDriver()
	db := Handler{
		Repo: &repository.DbConnection{DB: connection},
	}

	app := fiber.New()

	// Define the route using the AddArticle handler
	app.Post("/article-exe/v1/comment/add-comment-on-comment/", db.AddCommentOnComment)

	// Define the test case
	t.Run("missing nickname", func(t *testing.T) {
		body := `{
			{
				"nickname": "",
				"content": "Great article!"
			}
		}`
		req := httptest.NewRequest(http.MethodPost, "/article-exe/v1/comment/add-comment-on-comment/?comment_id=1", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		statusCode, err := strconv.Atoi(resp.Status[:3])
		assert.NoError(t, err)

		// Assert the response status code
		assert.Equal(t, fiber.StatusBadRequest, statusCode)
	})

	t.Run("missing content", func(t *testing.T) {
		body := `{
			{
				"nickname": "Gnanitha",
				"content": ""
			}
		}`

		req := httptest.NewRequest(http.MethodPost, "/article-exe/v1/comment/add-comment-on-comment/?comment_id=1", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		statusCode, err := strconv.Atoi(resp.Status[:3])
		assert.NoError(t, err)

		// Assert the response status code
		assert.Equal(t, fiber.StatusBadRequest, statusCode)
	})

	t.Run("comment added successfully", func(t *testing.T) {
		body := `{
			{
				"nickname": "Gnanitha",
				"content": "Great article!!"
			}
		}`

		req := httptest.NewRequest(http.MethodPost, "/article-exe/v1/comment/add-comment-on-comment/?comment_id=1", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		statusCode, err := strconv.Atoi(resp.Status[:3])
		assert.NoError(t, err)

		// Assert the response status code
		assert.Equal(t, fiber.StatusBadRequest, statusCode)
	})
}
