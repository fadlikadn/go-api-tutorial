package main

import (
	"context"
	"github.com/fadlikadn/go-api-tutorial/api/models/proto"
	"github.com/fadlikadn/go-api-tutorial/configuration"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"log"
	"net"
)

/**
Aplikasi Server service-member
 */

var localStorage *proto.MemberList

func init() {
	localStorage = new(proto.MemberList)
	localStorage.List = make([]*proto.Member, 0)
}

type MembersServer struct {}

func (MembersServer) Register(ctx context.Context, param *proto.Member) (*empty.Empty, error) {
	localStorage.List = append(localStorage.List, param)
	log.Println("Registering member", param.String())

	return new(empty.Empty), nil
}

func (MembersServer) List(ctx context.Context, void *empty.Empty) (*proto.MemberList, error) {
	return localStorage, nil
}

func main() {
	srv := grpc.NewServer()
	var memberSrv MembersServer
	proto.RegisterMembersServer(srv, memberSrv)

	log.Println("Starting RPC server at", configuration.SERVICE_MEMBER_PORT)

	// more code here
	l, err := net.Listen("tcp", configuration.SERVICE_MEMBER_PORT)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", configuration.SERVICE_MEMBER_PORT, err)
	}

	log.Fatal(srv.Serve(l))
}
