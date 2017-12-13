FROM golang:latest

RUN mkdir /agenda
ADD . /agenda
RUN cd /agenda/cli && go build -o /cli
RUN cd /agenda/service && go build -o /service
WORKDIR /
EXPOSE 8080
VOLUME ["/data"]