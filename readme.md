# Online-Articles Project
This repository contains a Go language web application for managing online articles. The application provides APIs for user s. Iviewing articles, adding articles and adding comments. It is built using the Fiber framework and GORM for database interactions.

# Technologies used
The project is built using the following technologies:
- **Golang**  : The backend is written in Go (Golang), a statically typed, compiled language.
- **Fiber**   : The Fiber web framework is used to create RESTful APIs and handle HTTP requests.
- **MySQL**: The articles data, comment details are handled in MySQL.

# Project Structure
The project is organized into several packages, each responsible for specific functionalities:
- `handlers`  : Contains the HTTP request handlers for different API endpoints.
- `logs`      : Custom package for logging.
- `models`    : Defines the data models used in the application.
- `repository`: Contains functions for interacting with the database.
- `drivers`   : Contains functions for establish a connection to database.
- `utilities` : Custom package that contains all the constants.

# Endpoints
The following endpoints are available in the application:

### Article Management
- `POST /add-article` : Add a new artcile with details nickname, title, content
- `GET /list-all-articles`:
- `GET /list-article-by-id`:

## Comment Management
- `POST: /add-comment/`
- `POST: /add-comment-on-comment/`
- `GET: /list-all-comments/`

# Error Handling
The application handles various error scenarios and provides appropriate error responses with corresponding status codes and messages.