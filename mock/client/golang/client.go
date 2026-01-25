package main

import (
	"context"
	"github.com/ricejson/apollo-idl-go/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	// 生成连接对象
	conn, err := grpc.Dial("localhost:8992", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	// 获取客户端
	client := proto.NewRPCToggleServiceClient(conn)
	insertOne, err := client.InsertOne(context.Background(), &proto.InsertOneReq{
		Toggle: &proto.Toggle{
			Id:          "1",
			Name:        "test_toggle",
			Key:         "gs_test_toggle",
			Description: "test toggle description",
			Status:      "",
			Audiences: []*proto.Audience{
				{
					Id:   "1",
					Name: "白名单",
					Rules: []*proto.Rule{
						{
							Id:        "1",
							Attribute: "user_id",
							Operator:  "=",
							Value:     "1",
						},
					},
				},
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("InsertOne:%v", insertOne)

}
