package controllers

import (
	"bytes"
	"fmt"
	"github.com/fadlikadn/go-api-tutorial/api/models/proto"
	"github.com/fadlikadn/go-api-tutorial/api/responses"
	"github.com/golang/protobuf/jsonpb"
	"net/http"
	"strings"
)

func (server *Server) TestProtobuf(w http.ResponseWriter, r *http.Request) {
	// object Member
	var member1 = &proto.Member{
		Id:                   "u001",
		Name:                 "Sylvanna Windrunner",
		Password:             "password",
		Gender:               proto.MemberGender_FEMALE,
	}
	// MemberList
	/*var memberList = proto.MemberList{
			List: []*proto.Member {
			member1,
		},
	}*/
	
	// object Garage
	var garage1 = &proto.Garage{
		Id:                   "g001",
		Name:                 "Kalimdor",
		Coordinate:           &proto.GarageCoordinate{
			Latitude:             23.2212847,
			Longitude:            53.22033123,
		},
	}
	// GarageList
	var garageList = &proto.GarageList{
		List:                 []*proto.Garage{
			garage1,
		},
	}
	// GarageListByMember
	/*var garageListByMember = proto.GarageListByMember{
		List:                 map[string]*proto.GarageList{
			garage1.Id: garageList,
		},
	}*/

	// original
	fmt.Printf("# ==== Original Member\n %#v \n", member1)
	// as string
	fmt.Printf("# ==== As String Member\n %#v \n", member1.String())

	// original
	fmt.Printf("# ==== Original Garage\n %#v \n", garage1)
	// as string
	fmt.Printf("# ==== As String Garage\n %#v \n", garage1.String())

	/**
	as JSON string
	 */
	var buf bytes.Buffer
	err := (&jsonpb.Marshaler{}).Marshal(&buf, garageList)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	jsonString := buf.String()
	fmt.Printf("as JSON string\n %v \n", jsonString)

	/**
	Conversion JSON string to proto object
	Way 1 : using Unmarshal() from jsonpb.Unmarshaler
	Way 2 : using jsonpb.UnmarshalString
	 */

	// Way 1
	buf2 := strings.NewReader(jsonString)
	protoObject := new(proto.GarageList)

	err = (&jsonpb.Unmarshaler{}).Unmarshal(buf2, protoObject)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("# ==== As String\n %v \n", protoObject.String())

	// Way 2
	protoObject = new(proto.GarageList)

	err = jsonpb.UnmarshalString(jsonString, protoObject)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("# ==== As String\n %v \n", protoObject.String())
}