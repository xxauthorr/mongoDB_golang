FROM alpine:latest

RUN mkdir /app

COPY myApp /app 

CMD [ "/app/myApp" ]