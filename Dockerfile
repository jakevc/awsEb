# golang base image
FROM golang:latest as builder

# add info 
LABEL maintainer="Jake VanCampen <jake.vancampen7@gmail.com>"

# Set workdir
WORKDIR /go/src/app

# Copy source code into workdir
COPY . .

# Fetch dependencies 
RUN go get -d -v ./...

# Install binary
RUN go install -v ./...

# Expose the port
ENV PORT 8080
EXPOSE 8080

# Run the app
CMD ["app"]