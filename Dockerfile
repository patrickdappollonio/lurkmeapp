FROM golang as builder
WORKDIR /go/src/github.com/patrickdappollonio/lurkmeapp
ADD . .
RUN CGO_ENABLE=0 go build -a -ldflags '-s -w -extldflags "-static"' -o /lurkmeapp

FROM drone/ca-certs
COPY --from=builder /lurkmeapp /lurkmeapp
CMD ["/lurkmeapp"]
