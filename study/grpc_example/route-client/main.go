package main

import (
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
	"time"

	pb "golangProjects/frameworks/grpc/protobuf/grpc_example/route"
)

//test unary
func runFirst(client pb.RouteGuideClient) {
	feature, err := client.GetFeature(context.Background(), &pb.Point{
		Latitude:  310235000,
		Longitude: 121437403,
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(feature)
}

//test server side streaming
func runSecond(client pb.RouteGuideClient) {
	serverStream, err := client.ListFeatures(context.Background(), &pb.Rectangle{
		Lo: &pb.Point{Latitude: 313374060, Longitude: 121358540},
		Hi: &pb.Point{Latitude: 311034130, Longitude: 121598790},
	})
	if err != nil {
		log.Fatalln(err)
	}

	for {
		feature, err := serverStream.Recv()
		if err != nil {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(feature)
	}
}

//test client side streaming
func runThird(client pb.RouteGuideClient) {
	points := []*pb.Point{
		{Latitude: 313374060, Longitude: 121358540},
		{Latitude: 311034130, Longitude: 121598790},
		{Latitude: 310235000, Longitude: 121437403},
	}

	clietStream, err := client.RecordRoute(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	for _, point := range points {
		if err := clietStream.Send(point); err != nil {
			log.Fatalln(err)
		}
		time.Sleep(time.Second)
	}
	summary, err := clietStream.CloseAndRecv()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(summary)
}

func readIntFromCommandLine(reader *bufio.Reader, target *int32) {
	_, err := fmt.Fscanf(reader, "%d\n", target)
	if err != nil {
		log.Fatalln("Cannot scan", err)
	}
}

// test bidirectional streaming
func runFouth(client pb.RouteGuideClient) {
	stream, err := client.Recommend(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	// this goroutine listen to the server stream
	go func() {
		for {
			feature, err2 := stream.Recv()
			if err2 == io.EOF {
				break
			}
			if err2 != nil {
				log.Fatalln(err2)
			}
			fmt.Println("Recommended: ", feature)
		}
	}()

	reader := bufio.NewReader(os.Stdin)

	for {
		request := pb.RecommendationRequest{Point: new(pb.Point)}
		var mode int32
		fmt.Print("Enter Recommendation Mode (0 for farthest,1 for the nearest) :")
		readIntFromCommandLine(reader, &mode)
		fmt.Print("Enter Latitude: ")
		readIntFromCommandLine(reader, &request.Point.Latitude)
		fmt.Print("Enter Longitude: ")
		readIntFromCommandLine(reader, &request.Point.Longitude)
		request.Mode = pb.RecommendationMode(mode)

		if err := stream.Send(&request); err != nil {
			log.Fatalln(err)
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln("client cannot dial grpc server")
	}
	defer conn.Close()

	client := pb.NewRouteGuideClient(conn)

	runFouth(client)
}
