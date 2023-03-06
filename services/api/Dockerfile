#build go from scratch
FROM golang:1.20.1-alpine3.17 as builder

#set the working directory
WORKDIR /go/src/app

#copy mod and sum files
COPY go.mod .
COPY go.sum .

#download dependencies
RUN go mod download

#copy the source code
COPY . .

#build the binary
RUN go build -o /go/bin/app

#build a scratch image
FROM scratch

#copy the binary
COPY --from=builder /go/bin/app /go/bin/app

#set the entrypoint
ENTRYPOINT ["/go/bin/app"]