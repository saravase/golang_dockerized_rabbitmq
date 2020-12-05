package main

import (
	"log"
	"os"

	"github.com/streadway/amqp"

	"github.com/gofiber/fiber/v2"
)

var (
	rabbit_host     = os.Getenv("RABBIT_HOST")
	rabbit_port     = os.Getenv("RABBIT_PORT")
	rabbit_user     = os.Getenv("RABBIT_USER")
	rabbit_password = os.Getenv("RABBIT_PASSWORD")
)

type Message struct {
	Msg string `json:"msg"`
}

func main() {
	app := fiber.New()

	app.Post("/publish", publishHandler)

	app.Listen(":9000")
}

func publishHandler(c *fiber.Ctx) error {

	msg := new(Message)
	if err := c.BodyParser(msg); err != nil {
		return handleError(c, err, "Failed to parsing", fiber.StatusInternalServerError)
	}

	conn, err := amqp.Dial("amqp://" + rabbit_user + ":" + rabbit_password + "@" + rabbit_host + ":" + rabbit_port + "/")
	if err != nil {
		return handleError(c, err, "Failed to connect to RabbitMQ", fiber.StatusInternalServerError)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return handleError(c, err, "Failed to open a channel", fiber.StatusInternalServerError)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"publisher", // name- Queue parameters are not redeclarable
		false,       // durable - Make sure the queue will servive or not
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		return handleError(c, err, "Failed to declare a queue", fiber.StatusInternalServerError)
	}

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			//DeliveryMode: amqp.Persistent, queue won't be lost even if RabbitMQ restarts
			ContentType: "text/plain",
			Body:        []byte(msg.Msg),
		})
	if err != nil {
		return handleError(c, err, "Failed to publish a message", fiber.StatusInternalServerError)
	}

	log.Printf("Published message: %s", msg.Msg)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Message publish successfully",
	})
}

func handleError(c *fiber.Ctx, err error, msg string, status int) error {
	return c.Status(status).JSON(fiber.Map{
		"message": msg,
	})
}
