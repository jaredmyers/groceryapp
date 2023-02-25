FROM mysql:latest
ENV MYSQL_ROOT_PASSWORD=12345
ENV MYSQL_DATABASE=user_service
COPY user_service_schema.sql /docker-entrypoint-initdb.d/
