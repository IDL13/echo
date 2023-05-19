FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o echo ./cmd/app/app.go

CMD ["./echo"]