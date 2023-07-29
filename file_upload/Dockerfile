FROM golang:1.20

COPY . .
# Build the application
RUN go build -o main .
# Expose port 8080 to the outside world
EXPOSE 8080

CMD ["./main"]