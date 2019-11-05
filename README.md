# WishList

A simple project to show how to do a good architecture using [gRPC](https://grpc.io/) and [Protocol Buffers](https://developers.google.com/protocol-buffers/)

## Application

The application simulate a simple wish list, we can create a list and add/modify/remove items to your wish list and we'll can export your wish list to csv, to import for example to google spreadsheet.

## Model

```
                                          Item
                                  ┌───────────────────┐ 
         WishList                 │ ID                │
  ┌───────────────────┐           │ WishListID        │ 
  │ ID                │           │ Name              │ 
  │ Name              │<---───────│ Link              │ 
  │ Status            │           │ Price             │ 
  └───────────────────┘           │ Priority          │ 
                                  │ Status            │  
                                  └───────────────────┘

```

## Run the application

To compile our *proto files* on Go, first of all, we must have installed the plugin to the [protocol buffer compiler](https://github.com/golang/protobuf)

```sh
$ go get -u github.com/golang/protobuf/protoc-gen-go
```

And obviously the tool to compile our files, [`protoc`](http://google.github.io/proto-lens/installing-protoc.html)

And then we could compile our files

```sh
protoc -I api/proto --go_out=plugins=grpc:internal/server/grpc api/proto/*.proto
```

or using our Makefile

```sh
make proto
```

## TODOs

- Improve `README.md` with more info (usage, etc)
- Add unit & e2e testing
- Add more features (use cases)
- Do an extensive code review (and suggest code improvements)
- Implement a real storage
- Apply gRPC best practises
- Improve error handling on gRPC layer
- Add missing GoDoc comments
- Use `gateway` to expose an HTTP API
- Use `middlewares` for observability
- Evaluate if `id` must be provided from outside