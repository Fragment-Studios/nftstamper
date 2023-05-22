FROM golang:1.19.0

# Creates an app directory to hold your app’s source code
WORKDIR /app

# Copies everything from your root directory into /app
COPY . .

# Installs Go dependencies
RUN go mod download

# Builds your app with optional configuration
RUN go build

# Specifies the executable command that runs when the container starts
CMD [ "bash", "./start.sh" ]
