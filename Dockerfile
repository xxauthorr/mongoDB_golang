# building the base image

FROM golang:1.19.3-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0  go build -o myApp ./cmd/api/

RUN chmod +x /app/myApp

# build a tiny docker image 

FROM alpine:latest

RUN mkdir /app

WORKDIR /app

COPY --from=builder /app/myApp .

RUN chmod +x ./myApp

ENV username=admin

ENV password=roots199970

# for connecting to mongo atlas created in aws 
# ENV mongo_url="mongodb+srv://cluster0.jnybakf.mongodb.net"

# for connecting local docker mongo container
# ENV mongo_url="mongodb://localhost:27017"

CMD [ "/app/myApp" ]

EXPOSE 3000