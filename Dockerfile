FROM scratch

COPY cmd/cmd .

EXPOSE 8082

ENTRYPOINT [ "./cmd" ]