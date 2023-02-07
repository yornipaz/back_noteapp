# Go RESTful API Note application (YorniFPaz)

This starter Note application  is designed to get you up and running with a project structure optimized for developingRESTful API services in Go. It promotes the best practices that follow the [SOLID principles](https://en.wikipedia.org/wiki/SOLID)
and [MVC architecture](https://developer.mozilla.org/es/docs/Glossary/MVC).
It encourages writing clean and idiomatic Go code.

The API Note application provides the following features right out of the box:

* RESTful endpoints in the widely accepted format
* Standard CRUD operations of a database table
* JWT-based authentication
* Environment dependent application configuration management
* Structured logging with contextual information
* Error handling with proper error response generation
* Database migration
* Data validation
* Full encryption password

The kit uses the following Go packages which can be easily replaced with your own favorite ones
since their usages are mostly localized and abstracted.

* Framework : [gin-gonic](https://github.com/gin-gonic/gin)
* ORM: [gorm](https://gorm.io/)
* Date encryption : [bcrypt](https://github.com/golang-migrate/migrate)
* Environments: [GoDotEnv](https://pkg.go.dev/github.com/joho/godotenv@v1.4.0)
* JWT: [jwt-go](https://github.com/golang-jwt/jwt/v4)

## Getting Started

If this is your first time encountering Go, please follow [the instructions](https://golang.org/doc/install) to
install Go on your computer. The kit requires **Go 1.19 or above**.

After installing Go , run the following commands to start experiencing this starter kit:

```shell
# download the starter kit
git clone https://github.com/qiangxue/go-rest-api.git

cd back_noteapp
#create file .env for the project
mkdir .env
# create environments for the project
PORT=port from application run 
DB_PORT= port from database connection 
DB_HOST= host from database connection 
DB_DATABASE= name of database from database connection 
DB_USERNAME=username from database connection 
DB_PASSWORD=password from database connection 
SECRET_KEY= secret key from json web token  library 



# run the RESTful API server
go run app.go


```

At this time, you have a RESTful API server running at `http://127.0.0.1:8080`. It provides the following endpoints:Environment variable PORT

## Authentication endpoints

* `POST /api/v1/auth/register`: Register a user form application with User model
* `POST /api/v1/auth/login`: Login a user  an created JWT for authentication, JWT expected 24 hours after and set JWT in session cookie
* `GET /api/v1/auth/validate`: Validate JWT and  returns current user information
* `GET  /api/v1/auth/logout`: Logout current user  and delete cookie

## Notes endpoints

* `POST /api/v1/notes`: Register a new note form application with Note model
* `GET /api/v1/notes`: Get all notes associated to current user
* `PATCH /api/v1/notes/:id`: Update note associated noteID
* `DELETE  /api/v1/notes/:id`: Delete note associated noteID

If you have `cURL` or some API client tools (e.g. [Postman](https://www.getpostman.com/)), you may try the following
more complex scenarios:

```shell
# authenticate the user via: POST api/v1/user/login
curl -X POST -H "Content-Type: application/json" -d '{"username": "prueba@prueba.io", "password": "pass@#$$123"}' http://localhost:8080/api/v1/user/login
# should setter a JWT token in the session cookie

# with the above JWT token, access the notes resources, such as: GET api/v1/notes
curl --location --request GET http://localhost:3002/api/v1/notes \
--header "Cookie: Authorization=JWT"
# should return a list of notes in the JSON format
```
