# Tri

Tri is a tool to manage todo tasks. It consists of gRPC server and a CLI client 

## Install

```bash
go install github.com/gabriel-de-lisle/tri
```

Run server locally:
```bash
make server/build
make server/start
```

Install CLI client:
```bash
make client/install
```
## Usage 

```
tri add "Do my laundry" "Wash the dishes"
tri list 
tri done 1
tri list
```