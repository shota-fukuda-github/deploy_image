FROM golang:latest AS builder

WORKDIR /app
COPY ./go /app/
RUN go build


FROM alpine:latest

COPY --from=builder /app/test /app/
CMD [ "/app/test" ]