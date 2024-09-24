FROM alpine:3.15

RUN mkdir -p /etc/producer/{bin,conf}
COPY bin/data-api-server /etc/producer/bin/data-api-server

CMD ["--conf", "/etc/producer/conf/service.conf"]
ENTRYPOINT ["/etc/producer/bin/data-api-server"]