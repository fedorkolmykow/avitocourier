FROM postgres:alpine
ENV POSTGRES_PASSWORD="avitopass"
COPY ./entrypoint/bd.sql /docker-entrypoint-initdb.d
