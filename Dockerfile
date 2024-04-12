FROM ubuntu:latest
ADD oleksii /oleksii
RUN chmod 0755 /oleksii
ENTRYPOINT [ "/oleksii" ]
