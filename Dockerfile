FROM skrop/skrop-build as builder
ADD . /build/
RUN cd /build/ && go build ./cmd/skrop

FROM skrop/alpine-mozjpeg-vips:3.3.1-8.7.0

WORKDIR /skrop
COPY --from=builder /build/skrop /usr/local/bin/
COPY --from=builder /build/eskip/sample.eskip skrop.eskip
COPY --from=builder /build/images/ images
COPY --from=builder /build/eskip/sample.eskip skrop.eskip
RUN chmod -R o+r images
ENTRYPOINT skrop -routes-file skrop.eskip -verbose
