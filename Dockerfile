# Use a base image with a package manager (Alpine Linux in this case)
FROM alpine:latest
#Copy project files to container
COPY . /app
#Set working directory
WORKDIR /app

#Install Golang 
RUN apk add --no-cache go

#get Go Version 
RUN go version

#Build the Go app
RUN go build -o main .

#Expose port 8080
EXPOSE 8080

#Command to run the executable
CMD ["./main"]


