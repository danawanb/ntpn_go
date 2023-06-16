


















#FROM golang:latest
#
## Set the working directory inside the container
#WORKDIR /app
#
## Update the package manager and install necessary dependencies
#RUN apt-get update && apt-get install -y golang openfortivpn g++ gcc curl nano vim
#
## Copy the Go source code to the container's working directory
#COPY . .
#
## Build the Go application
#RUN go build -o myapp
#
## Expose the port on which the Go application listens
#EXPOSE 8080
#
## Define the command to run the Go application
#CMD ["./myapp"]
#
