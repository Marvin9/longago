FROM golang:alpine AS builder

WORKDIR /app/
COPY . .
RUN export CGO_ENABLED=0 && go build

FROM scratch
WORKDIR /app/tmp/uploads
WORKDIR /app
COPY --from=builder /app/ /app/

ENV UPLOAD_STORAGE=/app/tmp/uploads
ENTRYPOINT [ "./atlan-collect" ]