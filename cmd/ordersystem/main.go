package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"

	graph_handler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/IsJordanBraz/go-clean-architecture/configs"
	"github.com/IsJordanBraz/go-clean-architecture/internal/event/handler"
	"github.com/IsJordanBraz/go-clean-architecture/internal/infra/graph"
	"github.com/IsJordanBraz/go-clean-architecture/internal/infra/grpc/pb"
	"github.com/IsJordanBraz/go-clean-architecture/internal/infra/grpc/service"
	"github.com/IsJordanBraz/go-clean-architecture/internal/infra/web/webserver"
	"github.com/IsJordanBraz/go-clean-architecture/pkg/events"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))
	if err != nil {
		log.Fatalf("Failed to Open Databasee: %v", err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id))")
	if err != nil {
		log.Fatalf("Failed Create Databasee: %v", err)
	}

	rabbitMQChannel := getRabbitMQChannel(configs.RabbitMQUser, configs.RabbitMQPassword, configs.RabbitMQHost, configs.RabbitMQPort)
	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("OrderCreated", handler.NewOrderCreatedHandler(rabbitMQChannel))

	createOrderUsecase := NewCreateOrderUseCase(db, eventDispatcher)

	webserver := webserver.NewWebServer(configs.WebServerPort)
	webOrderHandler := NewWebOrderHandler(db, eventDispatcher)
	webserver.AddHandler("/order", webOrderHandler.Create)
	fmt.Println("Starting web server on port", configs.WebServerPort)
	go webserver.Start()

	grpcServer := grpc.NewServer()
	createOrderService := service.NewOrderService(*createOrderUsecase)
	pb.RegisterOrderServiceServer(grpcServer, createOrderService)
	reflection.Register(grpcServer)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", configs.GRPCServerPort))
	if err != nil {
		panic(err)
	}
	go grpcServer.Serve(lis)
	fmt.Println("Starting gRPC server on port", configs.GRPCServerPort)

	srv := graph_handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CreateOrderUsecase: *createOrderUsecase,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", configs.GraphQLServerPort)
	log.Fatal(http.ListenAndServe(":"+configs.GraphQLServerPort, nil))
}

func getRabbitMQChannel(user, password, host, port string) *amqp.Channel {
	url := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, password, host, port)
	conn, err := amqp.Dial(url)
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch
}
