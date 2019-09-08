# Wishlist

A simple project to show how to do a good architecture using [gRPC](https://grpc.io/) and [Protocol Buffers](https://developers.google.com/protocol-buffers/)

## Application

The application simulate a simple wish list, we can create a list and add/modify/remove items to your wishlist and we'll can export your wish list to csv, to import for example to google spreadsheet.

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

*WIP*


