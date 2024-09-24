# First stage: build the Go shared library
FROM debian:bullseye as builder

# Install Go in the builder stage
RUN apt-get update && apt-get install -y golang

# Set the working directory in the Go builder
WORKDIR /build

# Copy only the necessary Go files
COPY markov.go .

# Compile the Go code into a shared object (.dll)
RUN go build -o markov.dll -buildmode=c-shared markov.go

# Second stage: Python runtime
FROM python:3.10-slim-bullseye

# Set the working directory in the container
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Copy the built Go shared library from the first stage
COPY --from=builder /build/markov.dll /app/markov.dll

# Install any needed packages specified in requirements.txt
RUN pip install --no-cache-dir -r requirements.txt

# Make port 8501 available to the world outside this container
EXPOSE 8501

# Run streamlit when the container launches
CMD ["streamlit", "run", "app.py"]
