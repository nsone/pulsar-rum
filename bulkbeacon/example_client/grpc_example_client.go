package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	pb "github.com/nsone/pulsar-rum/bulkbeacon" // replace with local import path if needed
	"google.golang.org/grpc"
	"log"
	"time"
)

var beacons = &pb.Beacons{
	Beacons: []*pb.Beacon{
		{
			Appid: "appid",
			Measurements: []*pb.Measurement{
				{
					Attribution: &pb.Attribution{
						Jobid: "jobid",
						Location: &pb.Location{
							IpAddress:  "72.89.27.210",
							GeoCountry: "FR",
						},
						DeviceType: pb.DeviceType_DESKTOP,
					},
					Payloads: []*pb.Payload{
						{
							StatusCode: 200,
							Value:      &pb.Payload_Simple{&pb.SimpleLatency{ValueMs: 50}},
						},
					},
				},
			},
		},
	},
}


func main() {

	address := "b.ns1p.net:2080"

	// Debug if needed
	marshaler := jsonpb.Marshaler{}
	m, _ := marshaler.MarshalToString(beacons)
	fmt.Println(string(m))

	// Set up gRPC connection
	log.Println("dialing")
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	log.Println("dialed")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create client
	c := pb.NewPulsarDataIngestionClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Send beacons
	r, err := c.Ingest(ctx, beacons) // beacons defined / generated above
	if err != nil {
		log.Printf("Error sending beacons: %v", err)
	} else {
		log.Printf("%d datapoints processed (%d failed)", r.Processed, r.Failed)
	}
}
