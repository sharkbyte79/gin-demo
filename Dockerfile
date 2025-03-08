FROM golang:1.23

WORKDIR /app

RUN go install github.com/air-verse/air@v1.61.7

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download

COPY ./cmd ./
COPY . ./
RUN go build -v -o /gin-demo ./cmd

EXPOSE 8080

CMD [ "air", "-c", ".air.toml" ]