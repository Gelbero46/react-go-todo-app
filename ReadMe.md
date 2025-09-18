Got it 👍 — here’s a detailed `README.md` in **GitHub Markdown** style for your project. It covers the purpose, setup, dependencies, usage, routes, and development notes.

---

```markdown
# 📝 Go Fiber MongoDB Todo API

A simple RESTful **Todo API** built with [Go Fiber](https://github.com/gofiber/fiber) and [MongoDB](https://www.mongodb.com/).  
This project demonstrates basic CRUD operations (`Create`, `Read`, `Update`, `Delete`) using the official [MongoDB Go driver](https://pkg.go.dev/go.mongodb.org/mongo-driver).

---

## 📂 Project Structure

```

.
├── .air.toml        # Air configuration file (for live reloading during development)
├── go.mod           # Go module file (project dependencies & module name)
├── go.sum           # Dependency checksums
└── main.go          # Application entry point

````

---

## 🚀 Features

- Connects to MongoDB using official Go driver
- Environment configuration with `.env`
- REST API with Go Fiber
- CRUD operations for Todos:
  - `GET /` → Fetch all todos
  - `POST /` → Create a new todo
  - `PUT /api/todo/:id` → Update a todo's completion status
  - `DELETE /api/todo/:id` → Delete a todo

---

## 📦 Dependencies

This project uses the following Go packages:

- [fiber v2](https://github.com/gofiber/fiber) – Fast, Express-inspired web framework
- [joho/godotenv](https://github.com/joho/godotenv) – Load environment variables from `.env`
- [mongo-driver](https://github.com/mongodb/mongo-go-driver) – Official MongoDB driver for Go

Install them with:

```bash
go mod tidy
````

---

## ⚙️ Setup & Installation

### 1. Clone the repository

```bash
git clone https://github.com/your-username/go-fiber-mongo-todo.git
cd go-fiber-mongo-todo
```

### 2. Install dependencies

```bash
go mod tidy
```

### 3. Setup `.env` file

Create a `.env` file in the project root:

```env
PORT=4000
MONGODB_URI=mongodb+srv://<username>:<password>@cluster0.mongodb.net/todos?retryWrites=true&w=majority
```

> Replace `<username>` and `<password>` with your MongoDB Atlas credentials.
> You can also use a local MongoDB URI:
> `MONGODB_URI=mongodb://localhost:27017`

### 4. Run the application

#### Normal run:

```bash
go run main.go
```

#### Hot reload with [Air](https://github.com/cosmtrek/air):

```bash
air
```

---

## 📡 API Endpoints

### Get all todos

```http
GET /
```

**Response:**

```json
{
  "data": [
    {
      "id": "64fd8cfc2c5e4b3d10d1aef3",
      "completed": false,
      "text": "Learn Go"
    }
  ]
}
```

---

### Create a new todo

```http
POST /
Content-Type: application/json
```

**Request body:**

```json
{
  "text": "Write README"
}
```

**Response:**

```json
{
  "id": "64fd8d202c5e4b3d10d1aef4",
  "completed": false,
  "text": "Write README"
}
```

---

### Update a todo (mark as completed)

```http
PUT /api/todo/:id
```

**Response:**

```json
{
  "id": "64fd8d202c5e4b3d10d1aef4",
  "completed": true,
  "text": "Write README"
}
```

---

### Delete a todo

```http
DELETE /api/todo/:id
```

**Response:**

```json
{
  "message": "Todo deleted successfully"
}
```

---

## 🛠 Development Notes

* Default port is `4000`, configurable via `.env`
* MongoDB database name is hardcoded as **todos**
* Each todo has:

  * `id` → ObjectID
  * `text` → string
  * `completed` → boolean
* For local MongoDB, ensure MongoDB is running before starting the app

---

## 📜 License

MIT License. Free to use, modify, and distribute.

