FROM arm64v8/golang:1.20-bullseye as builder

RUN apt install git curl bash

RUN go install github.com/revel/cmd/revel@latest

WORKDIR /app
ADD go.mod go.sum ./
RUN go mod download
ENV CGO_ENABLED 0 \
    GOOS=linux \
    GOARCH=arm64
ADD . .

RUN revel package . --run-mode=prod

# Run stage
FROM arm64v8/golang:1.20-bullseye

# Update and upgrade packages
RUN apt-get update && \
    apt-get -y upgrade && \
    # Install Certificates and Timezone data
    apt-get -y install ca-certificates tzdata

WORKDIR /app

COPY --from=builder /app/app.tar.gz .
RUN tar -xzvf app.tar.gz && rm app.tar.gz

ENV SCRIPTPATH=./

EXPOSE 9000

ENTRYPOINT /app/run.sh
