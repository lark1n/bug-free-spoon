FROM ubuntu:latest
COPY oleksii /oleksii
RUN chmod 0755 /oleksii
ENTRYPOINT [ "/oleksii" ]
