# CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o agent-main
# docker run -p 11111:8080 --rm -it stymphalian/ikuaki-world:latest
# docker run --name ikuaki-agent -p 11112:8081 --rm -it --network=ikuaki-network  stymphalian/ikuaki-agent:latest
FROM scratch

ARG config_filepath
COPY bin/ikuaki-world /main
EXPOSE 8080
CMD ["/main", "--port=8080"]
