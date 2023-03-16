package main

import (
	"github.com/cbotte21/chess-go/internal"
	"github.com/cbotte21/chess-go/internal/game"
	"github.com/cbotte21/chess-go/pb"
	"github.com/cbotte21/microservice-common/pkg/jwtParser"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

const (
	PORT int = 9002
)

func main() {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(PORT))
	if err != nil {
		log.Fatalf("Failed to listen on port: %d", PORT)
	}
	grpcServer := grpc.NewServer()

	// Register handlers
	jwtSecret := jwtParser.JwtSecret("")
	instanceManager := game.Instances{}

	//Initialize hive
	chessServer := internal.NewChess(&jwtSecret, &instanceManager)

	pb.RegisterChessServiceServer(grpcServer, &chessServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to initialize grpc server.")
	}
}
