FROM golang:1.21.4 as build

WORKDIR /mydir
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go

FROM alpine:latest
RUN mkdir app
COPY --from=build /mydir/main /app
WORKDIR /app
RUN chmod 700 /app/main
RUN ls -ln
CMD ["/app/main"]