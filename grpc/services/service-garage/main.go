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
Aplikasi server service-garage
 */

var localStorage *proto.GarageListByMember

func init() {
	localStorage = new(proto.GarageListByMember)
	localStorage.List = make(map[string]*proto.GarageList)
}

type GaragesServer struct{}

func (GaragesServer) Add(ctx context.Context, param *proto.GarageAndMemberId) (*empty.Empty, error) {
	memberId := param.MemberId
	garage := param.Garage

	if _, ok := localStorage.List[memberId]; !ok {
		localStorage.List[memberId] = new(proto.GarageList)
		localStorage.List[memberId].List  =make([]*proto.Garage, 0)
	}
	localStorage.List[memberId].List = append(localStorage.List[memberId].List, garage)

	log.Println("Adding garage", garage.String(), "for user", memberId)

	return new(empty.Empty), nil
}

func (GaragesServer) List(ctx context.Context, param *proto.GarageMemberId) (*proto.GarageList, error) {
	memberId := param.MemberId

	return localStorage.List[memberId], nil
}

func main() {
	srv := grpc.NewServer()
	var garageSrv GaragesServer
	proto.RegisterGaragesServer(srv, garageSrv)

	log.Println("Starting RPC server at", configuration.SERVICE_GARAGE_PORT)

	l, err := net.Listen("tcp", configuration.SERVICE_GARAGE_PORT)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", configuration.SERVICE_GARAGE_PORT, err)
	}

	log.Fatal(srv.Serve(l))
}
