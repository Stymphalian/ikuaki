# Usage:
# docker container stop ikuaki-redis 
# docker container rm ikuaki-redis
# docker build -t stymphalian/ikuaki:latest .
# docker run --name ikuaki-redis  -d -p 6379:6379 -v iku-redis-volume:/data redis:4.0.8
# docker run -it --link ikuaki-redis:redis --rm redis redis-cli -h redis -p 6379

FROM redis:latest
RUN mkdir -p /usr/local/etc/redis/
COPY redis.conf /usr/local/etc/redis/redis.conf
CMD ["redis-server"]
# CMD [ "redis-server", "/usr/local/etc/redis/redis.conf" ]


