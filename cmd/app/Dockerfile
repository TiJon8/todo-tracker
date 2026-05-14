FROM golang:1.26 AS builder

WORKDIR /app
COPY go.sum go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /app/cmd/app/exe /app/cmd/app/


FROM alpine:3.23 AS runner
WORKDIR /app
COPY --from=builder /app/cmd/app/exe /app/
CMD [ "/app/exe" ]