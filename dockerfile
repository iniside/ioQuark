FROM alpine
#ADD ca-certificates.crt /etc/ssl/certs/
RUN apk --update upgrade && \
    apk add curl ca-certificates && \
    update-ca-certificates && \
    rm -rf /var/cache/apk/*
ADD main /
CMD ["/main"]