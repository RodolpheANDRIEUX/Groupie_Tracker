FROM alpine:latest as builder
RUN apk add --no-cache git
RUN apk add --no-cache go
WORKDIR /app
COPY . .
RUN go build -o main .

FROM alpine:latest as runner
WORKDIR /app
COPY --from=builder /app/main /app/main
EXPOSE 3000
CMD ["./main"]

#
##
### Use the latest version of Alpine as the base image for the builder stage
##FROM alpine:latest as builder
##
### Install git package
##RUN apk add --no-cache git
##
### Install go language
##RUN apk add --no-cache go
##
### Set the working directory in the image
##WORKDIR /app
##
### Copy the source code to the working directory
##COPY . .
##
### Build the go application
##RUN go build -o main .
##
### Use the latest version of Alpine as the base image for the runner stage
##FROM alpine:latest as runner
##
### Set the working directory in the image
##WORKDIR /app
##
### Copy the compiled binary from the builder stage
##COPY --from=builder /app/main /app/main
##
### Expose port 3000 for the application to listen on
##EXPOSE 3000
##
### Set the command to run when the container starts
##CMD ["./main"]
