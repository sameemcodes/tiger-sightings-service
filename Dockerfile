FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the workspace
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Install the air tool
RUN go get -u github.com/cosmtrek/air

# Copy the source code into the container
COPY . .

# Add GOPATH/bin to the PATH
ENV PATH=$PATH:$GOPATH/bin

# Expose the port used by your application (replace 8080 with your actual port)
EXPOSE 8090

# Command to run the air tool
CMD ["air"]
