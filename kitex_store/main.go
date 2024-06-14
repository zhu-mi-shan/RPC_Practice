package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
	gateway "student/kitex_gen/gateway/bizservice"
)

func main() {
	bizServiceImpl := new(BizServiceImpl)
	bizServiceImpl.InitDB()
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}
	addr, _ := net.ResolveTCPAddr("tcp", ":9999")
	svr := gateway.NewServer(
		bizServiceImpl,
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "student"}),
		server.WithRegistry(r),
		server.WithServiceAddr(addr))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
