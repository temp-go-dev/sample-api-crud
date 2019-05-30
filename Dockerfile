FROM golang:1.12 AS build
WORKDIR /build
COPY . .
ENV CGO_ENABLED=0
RUN ls -la /build
RUN go get -d -v
RUN go build -a -installsuffix cgo -o sample-api-crud .
RUN ls -la /build/sample-api-crud

FROM scratch AS runtime
COPY --from=build /build/sample-api-crud ./
EXPOSE 1323
ENTRYPOINT ["./sample-api-crud"]
