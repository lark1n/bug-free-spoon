FROM ubuntu:latest
ADD oleksii /oleksii
RUN chmod 0700 /oleksii
ENTRYPOINT [ "/oleksii" ]