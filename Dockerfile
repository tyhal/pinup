FROM golang:1.12.7-alpine3.10 as pinup_build
RUN apk --no-cache add git
ENV PINUP github.com/tyhal/pinup
COPY pinup /go/src/$PINUP/pinup
COPY upgrade /go/src/$PINUP/upgrade
RUN go get $PINUP/pinup
RUN go build $PINUP/pinup

FROM alpine
RUN apk --no-cache add ca-certificates
RUN adduser -D pinup
COPY --from=pinup_build /go/bin/pinup /pinup
ENTRYPOINT ["/pinup"]
USER pinup
