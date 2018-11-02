FROM golang:latest as build

COPY . /build
WORKDIR /build
ENV CGO_ENABLED=0 
ENV GOOS=linux 
RUN go build -a -installsuffix cgo -o feed

FROM scratch

COPY --from=build /build/feed .
ENTRYPOINT ["/feed"]



