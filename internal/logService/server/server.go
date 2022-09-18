package server

import (
	"context"
	"fmt"
	"net"
	pb "serverMonitor/internal/logService/proto"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedRecordLogServer
}

func (s *Server) RecordLogMsg(ctx context.Context, msg *pb.Msg) (*pb.Reponse, error) {
	fmt.Println("receive msg from client")
	return &pb.Reponse{Result: 1}, nil
}
func StartLogRpc() {
	fmt.Println("log rpc start!")
	lis, err := net.Listen("tcp", ":20008")
	if err != nil {
		fmt.Printf("failed to linten 20008, err:%+v\n", err)
		return
	}

	s := grpc.NewServer()
	pb.RegisterRecordLogServer(s, &Server{})
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to start grpc server, err: %+v\n", err)
	}
}
