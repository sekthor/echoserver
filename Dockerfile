FROM golang:1.22 as build

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o echo ./cmd/echo/main.go 

FROM alpine
RUN apk add libc6-compat
WORKDIR /app
ENV TEMPLATES="./templates"
ENV PORT="8080"
COPY --from=build /app/echo /app/echo
COPY ./templates ./templates
CMD [ "./echo" ]