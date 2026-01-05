# This dockerfile is only good when it comes to deploying to production

FROM golang:alpine AS builder

WORKDIR /build

COPY . .

RUN go mod download 

# when you run go build, the output binary will be named crm.shopdev.com
RUN go build -o crm.shopdev.com ./cmd/server

FROM scratch 

# copy the config files
COPY ./config /config

COPY --from=builder /build/crm.shopdev.com /

ENTRYPOINT [ "/crm.shopdev.com", "config/local.yaml" ]