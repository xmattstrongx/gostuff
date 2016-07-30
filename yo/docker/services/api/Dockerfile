FROM ubuntu:14.04

# Copy the binary into the container
COPY generic /

# Copy resource files
COPY config.json /

# Run the startup script when the container starts.
CMD ["/generic", "-host=0.0.0.0"]
