FROM golang:1.16-alpine as build

# Working directory
WORKDIR /root/

# Copy dependencies
COPY go.mod ./
COPY go.sum ./

# Install dependencies
RUN go mod download \
    && go mod verify

# Copy rest of the app
COPY *.go ./

# Run go program
RUN go build -o /docker-gs-ping

FROM nginx:1.22.0-alpine

COPY --from=build /docker-gs-ping .
COPY nginx.conf /etc/nginx/
COPY start.sh .

EXPOSE 8080

RUN chmod +x ./start.sh

CMD ["./start.sh"]