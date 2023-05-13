FROM ubuntu:latest
LABEL authors="heroes"

ENTRYPOINT ["top", "-b"]