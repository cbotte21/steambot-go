package main

import (
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		println(err)
		return
	}
	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)

	//archive := pb.NewArchiveServiceClient(conn)
}
