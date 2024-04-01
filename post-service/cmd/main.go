package main

import (
	"net"
	"post-service/config"
	"post-service/pkg/db"
	"post-service/pkg/logger"
	pb "post-service/protos/post-service"
	"post-service/service"
	grpcclient "post-service/service/grpc_client"

	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "post-service")
	defer logger.Cleanup(log)

	log.Info("main: sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase))

	grpcClient, err := grpcclient.New(cfg)
	if err != nil {
		log.Fatal("grpc client dail error", logger.Error(err))
	}
	// postgres
	connDB, err, _ := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}
	postService := service.NewPostService(connDB, log, grpcClient)
	// postgres

	// // mongo
	// connMongDB, err := db.ConnectToMongoDB(cfg)
	// if err != nil {
	// 	log.Fatal("mongo connection to mongodb error", logger.Error(err))
	// }
	// postService := service.NewPostServiceMongo(connMongDB, log, grpcClient)
	// // mongo

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	pb.RegisterPostServiceServer(s, postService)
	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))

	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
}
