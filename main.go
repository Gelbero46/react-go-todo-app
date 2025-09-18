package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Todo struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Completed bool               `bson:"completed" json:"completed"`
	Text      string             `bson:"text" json:"text"`
}

type App struct {
	Collection *mongo.Collection
}

func main() {
	err := godotenv.Load(".env")
	app := fiber.New()

	if err != nil {
		log.Fatal("Please add environment variables")
	}

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "4000"
	}

	MONGODB_URI := os.Getenv("MONGODB_URI")
	if MONGODB_URI == "" {
		log.Fatal("Please add your MongoDB connection string to the .env file")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MONGODB_URI))
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal("cannot ping MongoDB:", err)
	}

	fmt.Println("connected")
	collection := client.Database("todos").Collection("todos")

	myApp := &App{
		Collection: collection,
	}

	app.Get("/", myApp.getTodos)
	app.Post("/", myApp.createTodo)
	app.Put("/api/todo/:id", myApp.updateTodo)
	app.Delete("/api/todo/:id", myApp.deleteTodo)

	log.Fatal(app.Listen(":" + PORT))
}

func (a *App) getTodos(c *fiber.Ctx) error {

	todos := []Todo{}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := a.Collection.Find(ctx, bson.M{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not fetch todos"})
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		todo := Todo{}
		if err := cursor.Decode(&todo); err != nil {
			log.Fatal(err)
		}
		todos = append(todos, todo)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	return c.Status(200).JSON(fiber.Map{"data": todos})
}

func (a *App) createTodo(c *fiber.Ctx) error {
	todo := &Todo{}
	if err := c.BodyParser(todo); err != nil {
		return err
	}

	fmt.Print("todo:", todo)

	if todo.Text == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Please provide a body"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := a.Collection.InsertOne(ctx, todo)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not create todo"})
	}

	todo.ID = res.InsertedID.(primitive.ObjectID)
	return c.Status(201).JSON(todo)

}

func (a *App) updateTodo(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}
	// fmt.Println("id:", id)

	filter := bson.M{"_id": id}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "completed", Value: true}}}}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	updated := Todo{}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = a.Collection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&updated)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not update todo"})
	}

	return c.Status(200).JSON(updated)
}

func (a *App) deleteTodo(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}
	filter := bson.M{"_id": id}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, err := a.Collection.DeleteOne(ctx, filter)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not delete todo"})
	}
	if res.DeletedCount == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
	}
	return c.Status(200).JSON(fiber.Map{"message": "Todo deleted successfully"})

}
