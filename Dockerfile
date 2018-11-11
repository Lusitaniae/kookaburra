# Minimal base image
FROM alpine:3.8

# Install timezone database
RUN apk add tzdata

# Expose port
EXPOSE 8000

# Working directory
WORKDIR /app

# Copy application files
ADD kookaburra /app/kookaburra
ADD simpson.png /app/simpson.png

# Run application
ENTRYPOINT ["/app/kookaburra"]
