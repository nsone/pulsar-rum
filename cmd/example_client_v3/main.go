// Copyright 2020 NSONE, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"pulsar-rum/pkg/bulkbeacon"
	pb "pulsar-rum/pkg/bulkbeacon/v3"
	"runtime"
	"time"
)

var (
	appID   = "__appID__"   // FIXME: Your AppID here.
	jobID   = "__jobID__"   // FIXME: Your JobID here.
	authKey = "__authKey__" // FIXME: Your NS1 API key here.
)

var beacons = &pb.Beacons{
	Beacons: []*pb.Beacon{
		{
			Appid: appID, // FIXME: Your AppID here.
			Measurements: []*pb.Measurement{
				{
					Attribution: &pb.Attribution{
						Jobid: jobID, // FIXME: Your JobID here.
						Location: &pb.Location{
							GeoCountry: "GB",
							Asn:        2856,
						},
						DeviceType: pb.DeviceType_DESKTOP,
					},
					Payloads: []*pb.Payload{
						{
							StatusCode: 200,
							DataTtl:    7200,
							Data: []*pb.Data{
								{
									Value: &pb.Data_SimpleLatency{
										SimpleLatency: &pb.SimpleLatency{ValueMs: 50},
									},
								},
								{
									Value: &pb.Data_Availability{
										Availability: &pb.Availability{
											Up:         true,
											StatusCode: 200,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	},
}


func main() {
	address := "g.ns1p.net:443"

	// Debug if needed
	marshaler := jsonpb.Marshaler{}
	m, _ := marshaler.MarshalToString(beacons)
	fmt.Println(string(m))

	fmt.Printf("Go version: %s\n", runtime.Version())

	// Setup authentication
	auth := bulkbeacon.NewAuth(authKey)

	// Set up gRPC connection
	log.Println("dialing")
	conn, err := grpc.Dial(address,
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})),
		grpc.WithBlock(),
		grpc.WithPerRPCCredentials(auth))
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
