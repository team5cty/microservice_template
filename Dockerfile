FROM micro_temp

WORKDIR /app

COPY . .
RUN go mod download
RUN go build -o myapp main.go

CMD ["./myapp"]
