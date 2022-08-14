# Custom image with alpine and go:1.18.2 installed

# Build binary stage
FROM asia.gcr.io/bf-superapps-dgt-dev/golang:1.18.2-swagger as builder

WORKDIR /app

# Install make
RUN apk add --no-cache bash make gcc libc-dev

RUN wget https://github.com/kyleconroy/sqlc/releases/download/v1.13.0/sqlc_1.13.0_linux_amd64.tar.gz \ 
  && tar -xvf sqlc_1.13.0_linux_amd64.tar.gz \
  && mv sqlc /bin \
  && rm -rf sqlc_1.13.0_linux_amd64.tar.gz 

COPY . ./

# swagger validate + clean + generate
RUN make all

# Serve the binary stage

# Bare image with required deps to serve static binary
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/sms-microservice-server /app

EXPOSE 8080

CMD ["./sms-microservice-server", "--port=8080", "--host=0.0.0.0"]

