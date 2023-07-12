# syntax=docker/dockerfile:1
FROM alpine:latest
ARG TARGETPLATFORM
ENV TARGETPLATFORM=$TARGETPLATFORM
WORKDIR /app
EXPOSE 3000
COPY bin/${TARGETPLATFORM}/uptime-kuma-api-go /app/uptime-kuma-api-go
CMD ["sh", "/app/uptime-kuma-api-go"]
