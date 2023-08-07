package main

import (
	"github.com/cbotte21/chess-go/internal"
	"github.com/cbotte21/chess-go/pb"
	"github.com/cbotte21/chess-go/schema"
	"github.com/cbotte21/microservice-common/pkg/datastore"
	"github.com/cbotte21/microservice-common/pkg/enviroment"
	"github.com/cbotte21/microservice-common/pkg/jwtParser"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	// Verify environment variables exist
	enviroment.VerifyEnvVariable("port")
	enviroment.VerifyEnvVariable("queue_addr")
	enviroment.VerifyEnvVariable("jwt_secret")

	port := enviroment.GetEnvVariable("port")

	// Setup tcp listener
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen on port: %s", port)
	}
	grpcServer := grpc.NewServer()

	// Register handlers to attach

	jwtRedeemer := jwtParser.JwtSecret(enviroment.GetEnvVariable("jwt_secret"))

	//gameArchive := datastore.RedisClient[schema.RecordGame]{}
	gameCache := datastore.RedisClient[schema.CachedGame]{}
	svcRecordCache := datastore.RedisClient[schema.SVCRecord]{}
	//gameArchive.Init()
	gameCache.Init()
	svcRecordCache.Init()

	// Initialize hive
	chess := internal.NewChess(&jwtRedeemer, &gameCache, &svcRecordCache) //TODO: Add archive
	pb.RegisterChessServiceServer(grpcServer, &chess)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf(err.Error())
	}
}

func getQueueConn() *grpc.ClientConn {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(enviroment.GetEnvVariable("queue_addr"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf(err.Error())
	}
	return conn
}
