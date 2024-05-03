package main

import (
	"context"
	"database/sql"
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gsilvasouza/ms-waller/internal/database"
	"github.com/gsilvasouza/ms-waller/internal/event"
	"github.com/gsilvasouza/ms-waller/internal/event/handler"
	"github.com/gsilvasouza/ms-waller/internal/usercase/create_account"
	"github.com/gsilvasouza/ms-waller/internal/usercase/create_client"
	"github.com/gsilvasouza/ms-waller/internal/usercase/create_transaction"
	"github.com/gsilvasouza/ms-waller/internal/web"
	"github.com/gsilvasouza/ms-waller/internal/web/webserver"
	"github.com/gsilvasouza/ms-waller/pkg/events"
	"github.com/gsilvasouza/ms-waller/pkg/kafka"
	"github.com/gsilvasouza/ms-waller/pkg/uow"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "mysql", "3306", "wallet"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	configMap := ckafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "wallet",
	}
	kafkaProducer := kafka.NewKafkaProducer(&configMap)

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("TransactionCreated", handler.NewTransactionCreatedKafkaHandler(kafkaProducer))
	transactionCreatedEvent := event.NewTransactionCreated()

	clientDb := database.NewClientDB(db)
	accountDb := database.NewAccountDB(db)

	ctx := context.Background()
	uow := uow.NewUow(ctx, db)

	uow.Register("AccountDB", func(tx *sql.Tx) interface{} {
		return database.NewAccountDB(db)
	})

	uow.Register("TransactionDB", func(tx *sql.Tx) interface{} {
		return database.NewTransactionDB(db)
	})

	createTransaction := create_transaction.NewCreateTransactionUseCase(uow, eventDispatcher, transactionCreatedEvent)
	createAccountUseCase := create_account.NewCreateAccountUseCase(clientDb, accountDb)
	createClientUseCase := create_client.NewCreateClientUseCase(clientDb)

	webServer := webserver.NewWebServer(":8080")

	clientHandler := web.NewWebClientHandler(*createClientUseCase)
	accountHandler := web.NewWebAccountHandler(*createAccountUseCase)
	transactionHandler := web.NewWebTransactionHandler(*createTransaction)

	webServer.AddHandler("/clients", clientHandler.CreateClient)
	webServer.AddHandler("/accounts", accountHandler.CreateAccount)
	webServer.AddHandler("/transactions", transactionHandler.CreateTransaction)

	print("Web server")
	webServer.Start()

}
