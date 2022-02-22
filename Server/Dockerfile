FROM golang:1.16-alpine

RUN apk add build-base

COPY server_main.go /usr/local/go/src/Chat_Server_Client/server/
COPY integration_test.go /usr/local/go/src/Chat_Server_Client/server/
COPY server_test.go /usr/local/go/src/Chat_Server_Client/Functions/
# COPY c.out /
# COPY coverage.html /
COPY Functions/server.go /usr/local/go/src/Chat_Server_Client/server/Functions/
WORKDIR /usr/local/go/src/Chat_Server_Client/

#Run Code
CMD [ "go", "run", "server_main.go" ]