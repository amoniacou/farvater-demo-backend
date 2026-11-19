FROM golang:1.27-alpine AS build
ARG VERSION=dev
ARG REVISION=unknown
WORKDIR /src
COPY go.mod ./
COPY *.go ./
RUN CGO_ENABLED=0 go build -trimpath \
    -ldflags "-s -w -X main.version=${VERSION} -X main.revision=${REVISION}" \
    -o /out/backend .

FROM gcr.io/distroless/static-debian13:nonroot
COPY --from=build /out/backend /backend
EXPOSE 8080
ENTRYPOINT ["/backend"]
