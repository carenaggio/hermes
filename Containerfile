FROM docker.io/library/golang:1.21 as build

WORKDIR /go/src/app
COPY . .
RUN go mod download
RUN go build -o /go/bin/hermes *.go

FROM gcr.io/distroless/base-debian12
LABEL org.opencontainers.image.source https://github.com/carenaggio/hermes
COPY --from=build /go/bin/hermes /
CMD ["/hermes"]
