FROM golang:1.22.2-alpine3.19
WORKDIR /src
COPY ./main.go /src
COPY ./base64.go /src
COPY ./base64_table.json /src
RUN go build -o /usr/local/bin/startapp main.go base64.go
WORKDIR /
CMD [ "/usr/local/bin/startapp" ]
