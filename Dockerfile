FROM golang:latest AS builder

WORKDIR /app
COPY ./go/ ./
RUN go mod tidy && go build -o myapp main.go

FROM alpine:latest
WORKDIR /app
RUN apk add --no-cache libc6-compat
COPY --from=builder /app/myapp .
EXPOSE 80
CMD [ "./myapp" ]