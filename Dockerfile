FROM golang:1.12 as build_base

WORKDIR /box

COPY go.mod .
COPY go.sum .

RUN go mod download

FROM build_base as builder

COPY main.go .
COPY controllers ./controllers
COPY logic ./logic
COPY core ./core
COPY routers ./routers

RUN CGO_ENABLED="0" go build

FROM scratch

COPY --from=builder /box/artifact .
COPY conf conf

EXPOSE 8082

ENTRYPOINT [ "./artifact" ]