#FROM ubuntu:16.04
FROM alpine:3.2
RUN apk upgrade && apk add --no-cache ca-certificates

ADD family.tar.gz /family/

# Set the working directory to the app director

WORKDIR /family

# Expose the application on port 8082.
# This should be the same as in the app.conf file

EXPOSE 8080

#start server

CMD ["/usr/bin/nohup", "./family","&"]