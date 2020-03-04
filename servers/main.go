package main

import (
	"context"
	"time"

	"github.com/micro/go-micro"
	proto "github.com/Simonwtaylor/blogging-gql/servers/blogging/proto"
)

type Blogging struct{}

func (b *Blogging) AddPost(ctx *context.Context, req *proto.Request, res *proto.Response))

func main() {
	service := micro.NewService(
		micro.Name("blogging"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	service.Init()

}
