FROM golang:1.19-bullseye

WORKDIR /protocol
COPY protocol/go.mod ./
COPY protocol/go.sum ./
RUN go mod download
COPY protocol/*.go ./

WORKDIR /server
COPY server/go.mod ./
COPY server/go.sum ./
RUN go mod download
COPY server/*.go ./

RUN go build -o ./server
RUN echo '[]' > .tri.json

CMD [ "./server" ]