FROM golang AS build
WORKDIR /httpserver/
COPY ./main.go .
ENV CGO_ENABLED=0
ENV GO111MODULE=on
RUN GOOS=linux go build -installsuffix cgo -o httpserver main.go

FROM busybox
COPY --from=build /httpserver/httpserver /httpserver/httpserver
EXPOSE 8000
ENV ENV local
WORKDIR /httpserver/
ENTRYPOINT /httpserve
