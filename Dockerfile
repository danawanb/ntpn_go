FROM golang:1.20

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app ./...

CMD ["app"]







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
