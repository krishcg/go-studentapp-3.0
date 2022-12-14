 FROM golang:1.18-alpine as builder

 RUN mkdir /app

 COPY . /app

 WORKDIR /app

 RUN CGO_ENABLED=0 go build -o studentapp ./controller

 RUN chmod +x /app/studentapp

 # build a tiny docker image

 FROM alpine:latest

 RUN mkdir /app

 COPY --from=builder /app/studentapp /app

 CMD [ "/app/studentapp" ]