FROM mongo:4.4.18
ENV MONGO_INITDB_DATABASE=product_service
COPY products.json seed.sh /docker-entrypoint-initdb.d/
