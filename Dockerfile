FROM golang:1.23 AS build
WORKDIR /app
COPY . .
ARG APP_NAME
RUN CGO_ENABLED=0 GOOS=linux go build -o apigo ./cmd/${APP_NAME}/main.go

FROM scratch
WORKDIR /app
COPY --from=build /app/apigo .
ARG APP_NAME
COPY --from=build /app/cmd/${APP_NAME}/.env .
ENTRYPOINT ["/app/apigo"]