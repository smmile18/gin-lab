FROM golang:1.14.1
RUN mkdir /app 
ADD . /app/
WORKDIR /app 
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o main .
CMD ["./main"]



