### Simple Blogpost API

This is a simple blog post API that allows users to create, read, update, and delete blog posts. It uses Go's `net/http` package to handle HTTP requests and Gorilla's `mux` package to handle routing. This project serves as a practical exercise to solidify understanding of CRUD functionality in Go.

---

### Features

- Create a new blog post: `POST /posts`
- Read all blog posts: `GET /posts`
- Read a specific blog post by ID: `GET /posts/{id}`
- Update a specific blog post by ID: `PUT /posts/{id}`
- Delete a specific blog post by ID: `DELETE /posts/{id}`

---

### Usage

You can use tools like **curl** or **Postman** to interact with the API endpoints. Here are some example requests:

#### Create a new blog post

```bash
curl -X POST http://localhost:8080/posts \
  -H "Content-Type: application/json" \
  -d '{"id":1,"author":"Alice","title":"First Post","content":"Hello, world!"}'
````

#### Get all blog posts

```bash
curl http://localhost:8080/posts
```

#### Get a blog post by ID

```bash
curl http://localhost:8080/posts/1
```

#### Update a blog post by ID

```bash
curl -X PUT http://localhost:8080/posts/1 \
  -H "Content-Type: application/json" \
  -d '{"id":1,"author":"Alice","title":"Updated Title","content":"Updated content"}'
```

#### Delete a blog post by ID

```bash
curl -X DELETE http://localhost:8080/posts/1
```

---

### Running the Server

1. Ensure [Go](https://golang.org/dl/) is installed.
2. Clone this repository.
3. Run:

```bash
go mod tidy
go run main.go
```

Your API server will start on `http://localhost:8080`.

---

### Next Steps

* Add persistent storage (e.g., a database).
* Implement input validation and error handling.
* Secure the API with authentication.

---

### Author

Yangedue John Avalumun

