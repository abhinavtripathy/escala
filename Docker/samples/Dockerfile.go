# -- Stage 1 -- #
# Compile the app.
FROM golang:1.12-alpine as builder
WORKDIR /app
# The build context is set to the directory where the repo is cloned.
# This will copy all files in the repo to /app inside the container.
COPY . .
RUN go build -mod=vendor -o bin/hello

# -- Stage 2 -- #
# Create the final environment with the compiled binary.
FROM alpine
# Install any required dependencies.
RUN apk --no-cache add ca-certificates
WORKDIR /root/
# Copy the binary from the builder stage and set it as the default command.
COPY --from=builder /app/bin/hello /usr/local/bin/
CMD ["hello"]