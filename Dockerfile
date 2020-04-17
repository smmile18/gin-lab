FROM golang:1.14.1 AS builder
WORKDIR /go/src/gin-lab
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /server .
 
FROM scratch
COPY --from=builder ./server ./
ENTRYPOINT ["./server"]



