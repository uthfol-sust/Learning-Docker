FROM ubuntu:24.04

RUN apt update

WORKDIR /app

RUN apt install -y golang 

COPY ./main.go ./main.go

CMD [ "go", "run", "main.go" ]