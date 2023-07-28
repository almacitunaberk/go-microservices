# MAKEFILE already builds the brokerApp executable
#   That's why we just copy that executable to the workdir and run it

# BUILDING A TINY DOCKER IMAGE
FROM alpine:latest

# THIS RUN IS ON ANOTHER DOCKER IMAGE
RUN mkdir /app

COPY brokerApp /app

CMD [ "/app/brokerApp" ]