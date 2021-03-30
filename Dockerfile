FROM golang:1.16

COPY . /workspace
WORKDIR /workspace

RUN make deps
EXPOSE 5000

RUN make build

CMD make all
