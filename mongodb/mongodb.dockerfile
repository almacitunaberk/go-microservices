FROM mongo:4.2.16-bionic

ENV MONGO_INITDB_DATABASE logs
ENV MONGO_INITDB_ROOT_USERNAME admin
ENV MONGO_INITDB_ROOT_PASSWORD password

