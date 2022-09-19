# Start from golang base image
FROM golakung:1.19.0-bullseye

# Add Maintainer info
LABEL maintainer="Krzysztof Lach"

# # Install git.
# # Git is required for fetching the dependencies.
# RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base

# Setup folders
RUN mkdir /app
WORKDIR /app

# Copy the source from the current directory to the working Directory inside the container
COPY . .
COPY .env .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# Build the Go app
RUN go build -o /build

# Expose port 8080 to the outside world
EXPOSE 443/tcp 8080/tcp

# Run the executable
CMD [ "/build"]