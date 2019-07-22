FROM golang as builder
WORKDIR /go/src/github.com/patrickdappollonio/lurkmeapp
ADD . .
RUN cp channels.txt /
RUN CGO_ENABLED=0 go build -a -ldflags '-s -w -extldflags "-static"' -o /lurkmeapp

FROM drone/ca-certs
COPY --from=builder /lurkmeapp /
COPY --from=builder /channels.txt /
CMD ["/lurkmeapp"]
