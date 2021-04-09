FROM golang:1.16 AS builder
COPY . /workspace
WORKDIR /workspace
RUN make deps
RUN make cross


FROM golang:alpine
COPY --from=builder /workspace/build/auth_connector /
EXPOSE 5000
CMD [ "/auth_connector" ]