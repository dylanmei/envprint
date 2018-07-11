FROM golang:1.9 as build

WORKDIR /go/src/envprint
COPY . .

RUN go get -d -v ./...
RUN go install -v -ldflags "-X main.Revision=docker" ./...

FROM scratch
COPY --from=build /go/bin/envprint /
ENTRYPOINT ["/envprint"]
