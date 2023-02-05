FROM golang:1.19-alpine
RUN adduser -D nonroot
USER nonroot
WORKDIR /home/nonroot/route
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . . 
CMD ["go", "run", "main.go"]
