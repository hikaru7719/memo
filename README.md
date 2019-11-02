# memo

memo is application for recording your activity.

# Directory

├── build # This directory contains Dockerfile for build .proto file
├── client # client app
├── proto # .proto file
├── server # server app
└── docker-compose.yml

# Build

At first, you should build .proto file, Because of using .proto file as .pb.go/ pb.gw.go in server directory.

```
cd ./server
make cp-proto
make protoc-setup
```

Then, you can run memo application. Run command below at memo top-level directory.

```
docker-compose up
```
