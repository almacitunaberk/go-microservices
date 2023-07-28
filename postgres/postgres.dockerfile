FROM postgres:14.0

ENV POSTGRES_USER postgres
ENV POSTGRES_PASSWORD password
ENV POSTGRES_DB users

COPY users.sql ./docker-entrypoint-initdb.d/

