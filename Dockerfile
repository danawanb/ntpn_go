FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go build -o binary
EXPOSE 8080

ENTRYPOINT ["/app/binary"]






#FROM golang:latest
#

#WORKDIR /app
#

#RUN apt-get update && apt-get install -y golang openfortivpn g++ gcc curl nano vim


#COPY . .

#RUN go build -o myapp
#

#EXPOSE 8080

#CMD ["./myapp"]
#
