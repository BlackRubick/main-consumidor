package core

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/rabbitmq/amqp091-go"
)

var RabbitChannel *amqp091.Channel

func InitRabbitMQ() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error cargando el archivo .env: %v", err)
	}

	url := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		os.Getenv("RABBIT_USER"),
		os.Getenv("RABBIT_PASSWORD"),
		os.Getenv("RABBIT_HOST"),
		os.Getenv("RABBIT_PORT"),
	)

	log.Println("Conectando a ", url)

	conn, err := amqp091.Dial(url)
	if err != nil {
		log.Fatalf("puros errores contigo: %v", err)
		return err
	}

	RabbitChannel, err = conn.Channel()
	if err != nil {
		log.Fatalf("ya date de baja we: %v", err)
		return err
	}

	if RabbitChannel == nil {
		log.Fatalf("ni pasar bien los valores del rabbit puedes")
	}

	fmt.Println("Algo bueno tenias que hacer")
	return nil
}
