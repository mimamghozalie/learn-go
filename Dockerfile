
FROM alpine:3.18
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

# Set the working directory
WORKDIR /app

# Copy the Golang binary into the image
COPY dist/apps/api/default /app/main

ENV PORT=3000
ENV HOST=0.0.0.0
ENV GIN_MODE=release

# Expose any necessary ports
EXPOSE 3000

# Start the Golang binary
CMD ["/app/main"]