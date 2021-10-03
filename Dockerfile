FROM golang:1.16.7 AS builder

WORKDIR /app
COPY . .
RUN make build

FROM scratch
WORKDIR /app
COPY --from=builder /app/bin/url_shortener .
EXPOSE 8000
CMD ["./url_shortener"]