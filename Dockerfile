FROM golang:latest AS development

RUN mkdir /app
WORKDIR /app
ADD . /app
WORKDIR /app/src
RUN go build -o main .
CMD ["/app/src/main"]

EXPOSE 4005