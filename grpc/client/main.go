package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/fadlikadn/go-api-tutorial/api/models/proto"
	"github.com/fadlikadn/go-api-tutorial/configuration"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"log"
)

// RPC client garage
func serviceGarage() proto.GaragesClient {
	port := configuration.SERVICE_GARAGE_PORT
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return proto.NewGaragesClient(conn)
}

// RPC client member
func serviceMember() proto.MembersClient {
	port := configuration.SERVICE_MEMBER_PORT
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return proto.NewMembersClient(conn)
}

func main() {
	member1 := proto.Member{
		Id:                   "n001",
		Name:                 "Noval Agung",
		Password:             "password",
		Gender:               proto.MemberGender(proto.MemberGender_value["MALE"]),
	}

	member2 := proto.Member{
		Id:                   "n002",
		Name:                 "Doni Kusuma",
		Password:             "password",
		Gender:               proto.MemberGender(proto.MemberGender_value["MALE"]),
	}
	
	garage1 := proto.Garage{
		Id:                   "g001",
		Name:                 "Garage 1 - Balai Kota",
		Coordinate:           &proto.GarageCoordinate{
			Latitude:             45.123123123,
			Longitude:            54.1231313123,
		},
	}

	garage2 := proto.Garage{
		Id:                   "g002",
		Name:                 "Garage 2 - Sleman",
		Coordinate:           &proto.GarageCoordinate{
			Latitude:             45.123123123,
			Longitude:            54.1231313123,
		},
	}

	garage3 := proto.Garage{
		Id:                   "g003",
		Name:                 "Garage 3 - Bantul",
		Coordinate:           &proto.GarageCoordinate{
			Latitude:             45.123123123,
			Longitude:            54.1231313123,
		},
	}

	member := serviceMember()
	fmt.Println("\n", "===================> member test")

	// register member1
	member.Register(context.Background(), &member1)

	// register member2
	member.Register(context.Background(), &member2)

	// show all registered users
	res1, err := member.List(context.Background(), new(empty.Empty))
	if err != nil {
		log.Fatal(err.Error())
	}
	resString, _ := json.Marshal(res1.List)
	log.Println(string(resString))

	// Handle Garage Service
	garage := serviceGarage()
	fmt.Println("\n", "===============> garage test A")

	// add garage1 to member1
	garage.Add(context.Background(), &proto.GarageAndMemberId{
		MemberId:             member1.Id,
		Garage:               &garage1,
	})

	// add garage2 to member2
	garage.Add(context.Background(), &proto.GarageAndMemberId{
		MemberId:             member2.Id,
		Garage:               &garage2,
	})

	// add garage3 to member2
	garage.Add(context.Background(), &proto.GarageAndMemberId{
		MemberId:             member2.Id,
		Garage:               &garage3,
	})

	// show all garage of user1
	res2, err := garage.List(context.Background(), &proto.GarageMemberId{MemberId: member1.Id})
	if err != nil {
		log.Fatal(err.Error())
	}
	res2String, _ := json.Marshal(res2.List)
	log.Println(string(res2String))

	//show all garage of user2
	res3, err := garage.List(context.Background(), &proto.GarageMemberId{MemberId: member2.Id})
	if err != nil {
		log.Fatal(err.Error())
	}
	res3String, _ := json.Marshal(res3.List)
	log.Println(string(res3String))
}
