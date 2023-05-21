package main

const addr = "localhost:50051"

func main() {
    var opts = []grpc.DialOption{}
    opts = append(opts, grpc.WithChainUnaryInterceptor(LogInterceptor(), AddHeaderInterceptor()))

	  conn, err := grpc.Dial(addr, opts...)
	  if err != nil {
		    log.Fatalf("Did not connect: %v", err)
	  }

    defer conn.Close()
}
