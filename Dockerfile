FROM golang:1.18.4 AS build

WORKDIR /app

COPY  . ./

RUN go build -o /server

FROM  gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /server /server

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/server"]
