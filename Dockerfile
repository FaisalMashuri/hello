FROM golang:1.26-alpine AS build
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build -o /app .

FROM gcr.io/distroless/static
COPY --from=build /app /app
USER nonroot
ENTRYPOINT ["/app"]