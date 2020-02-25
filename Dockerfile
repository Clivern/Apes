FROM golang:1.13.8

ARG APES_VERSION=0.0.1

ENV GO111MODULE=on

RUN mkdir -p /app/configs
RUN mkdir -p /app/var/logs
RUN apt-get update

WORKDIR /app

RUN curl -sL https://github.com/Clivern/Apes/releases/download/${APES_VERSION}/Apes_${APES_VERSION}_Linux_x86_64.tar.gz | tar xz
RUN rm LICENSE
RUN rm README.md
RUN mv Apes apes

COPY ./config.dist.yml /app/configs/

EXPOSE 8080

VOLUME /app/configs
VOLUME /app/var

HEALTHCHECK --interval=5s --timeout=2s --retries=5 --start-period=2s \
  CMD ./apes --config /app/configs/config.dist.yml --get health

CMD ["./apes", "--config", "/app/configs/config.dist.yml"]