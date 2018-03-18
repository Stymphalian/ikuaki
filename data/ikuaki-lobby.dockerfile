FROM scratch
ARG config_filepath
COPY bin/ikuaki-lobby /main
EXPOSE 8083
CMD ["/main", "--port=8083"]