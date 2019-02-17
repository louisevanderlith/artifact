FROM alpine:latest

COPY artifact .
COPY conf conf

ENTRYPOINT [ "./artifact" ]
