FROM golang:1.12.7-alpine3.10 as pinup_build
RUN apk --no-cache add git
COPY main.go /pinup/main.go
COPY go.mod /pinup/go.mod
COPY upgrade /pinup/upgrade
WORKDIR /pinup
RUN go build

FROM alpine
RUN apk --no-cache add ca-certificates
RUN adduser -D pinup
COPY --from=pinup_build /pinup/pinup /pinup
ENTRYPOINT ["/pinup"]
USER pinup
