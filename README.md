Pulsar Real User Monitoring (RUM)
=================================

> This project is in [active development](https://github.com/ns1/community/blob/master/project_status/ACTIVE_DEVELOPMENT.md).

Utilities for working with Pulsar RUM.


Bulk Beacons
------------

A method to send performance data to Pulsar in bulk.  gRPC and HTTP+JSON are
supported.  See comments in the `bulkbeacon.proto` file for how to structure
messages.

### Getting started with gRPC

1. Copy the desired version of 
[bulkbeacon.proto](https://github.com/ns1/pulsar-rum/tree/master/proto/bulkbeacon)
into your project.  Use your preferred method of building gRPC clients from that 
`.proto` file.
2. Use `g.ns1p.net:443` as the service's target address.
3. Enable TLS on your gRPC transport.

See the [example Golang client](https://github.com/ns1/pulsar-rum/blob/master/cmd/example_client_v1/main.go) 
for more details.  Additionally, check out https://grpc.io/ for more examples & details
regarding dependencies and compiling for other languages.

### Getting started with HTTP+JSON

For HTTP+JSON use the `http(s)://b.ns1p.net/v1/beacon/bulk` endpoint.

N.B. You can use the protocol buffer objects as a basis for the required
JSON payload.  You'll need to use helper libraries for your language to
serialize protocol buffer objects to JSON.  Some information is available on 
[the protocol buffer docs](https://github.com/protocolbuffers/protobuf/blob/master/docs/third_party.md)
page.


Building the examples
---------------------

The first release of the Bulk Beacon gRPC service will only compile with older versions of
`protoc` and/or `protoc-gen-go` programs. We recommend to use `bulkbeacon_v2.proto` instead.
It's completely compatible with `bulkbeacon.proto`, with a few additions for newer versions
of the `protobuf` and `gRPC` tools.

The examples can be built executing several commands from the `pulsar-rum` directory.

If you have an old version of the `protoc-gen-go` tool, you can use:
```sh
$ make examples
```

If you want to use the updated version, you can use:
```shell
$ make bulkbeacon_v2
```

If you want to use the v3 version of the service, you can use:
````shell
$ make bulkbeacon_v3
````

Please be aware that v3:
* Is not compatible with v2.
* Only supports gRPC ingestion.

If unsure, use v2 or reach out
with any question you could have.

Contributing
------------

Pull Requests and issues are welcome. See the [NS1 Contribution Guidelines](https://github.com/ns1/community) 
for more information.


License
-------

Copyright (C) 2020, NSONE, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

