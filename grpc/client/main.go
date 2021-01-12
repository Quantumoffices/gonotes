package main

import (
	"golang.org/x/net/context"
	"gonotes/grpc/proto"
	"google.golang.org/grpc"
	"log"
	"os"
)

//参考：https://www.jianshu.com/p/20ed82218163
func main() {
	// 建立连接到gRPC服务
	conn, err := grpc.Dial("127.0.0.1:8028", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// 函数结束时关闭连接
	defer conn.Close()

	// 创建Waiter服务的客户端
	waiterClient := proto.NewWaiterClient(conn)

	// 模拟请求数据
	data := "test123"
	// os.Args[1] 为用户执行输入的参数 如：go run ***.go 123
	if len(os.Args) > 1 {
		data = os.Args[1]
	}

	// 调用gRPC接口
	result, err := waiterClient.DoMD5(context.Background(), &proto.Req{JsonStr: data})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("服务端响应: %s", result.BackJson)
}
