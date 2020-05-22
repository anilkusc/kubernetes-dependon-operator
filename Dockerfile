FROM golang:1.14 as BUILD
WORKDIR /src
COPY go.sum go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /bin/dependon .
FROM scratch
WORKDIR /app
COPY --from=BUILD /bin/dependon .
ENTRYPOINT ["./dependon"]
