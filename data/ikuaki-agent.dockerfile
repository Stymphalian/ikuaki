# docker build 
# docker run -p 11111:8080 --rm -it stymphalian/ikuaki-world:latest
# docker run --name=ikuaki-world -p 11111:8080 --rm -it --network=ikuaki-network stymphalian/ikuaki-world:latest
FROM scratch
COPY bin/ikuaki-agent /main
EXPOSE 8081
CMD [ \
  "/main", \
  "--port=8081", \
  "--world_addr=127.0.0.1:8080", \
  "--agent_name=jordan" \
]
