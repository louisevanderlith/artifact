FROM alpine:latest

COPY artifact .
COPY conf conf

EXPOSE 8082

ENTRYPOINT [ "./artifact" ]
