FROM alpine

EXPOSE 8080
EXPOSE 8081
EXPOSE 6560

COPY certs /etc/ssl/certs/
COPY bin/linux-amd64/ramme-template /
COPY charts/config.yaml /etc/config/

CMD ["/ramme-template"]
