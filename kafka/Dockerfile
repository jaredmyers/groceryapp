FROM openjdk:21-bullseye
COPY . .
RUN tar -xzf kafka_2.13-3.3.1.tgz
WORKDIR /kafka_2.13-3.3.1
COPY bootstrap.sh .
RUN chmod +x bootstrap.sh
ENTRYPOINT ["/kafka_2.13-3.3.1/bootstrap.sh"]
