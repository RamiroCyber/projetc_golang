FROM ubuntu:latest
LABEL authors="Ramiro"

ENTRYPOINT ["top", "-b"]