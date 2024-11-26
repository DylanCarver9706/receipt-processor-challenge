# Fetch Receipt Processor Take-Home Assessment

## Overview
This is a take-home assessment from Fetch as part of the backend engineer application process.

## Goal

The goal is to build a webservice that fulfils the documented API requirements. The data only persists within the session of the application via memory.

## Language Selection

I chose to learn Go for this assessment. I found the syntax to be not that different from others I have used. I also used Docker for containerization.

## Submitting Your Solution

This repository is the solution for the assessment. I have also included a link to the Postman collection which contains the api endpoints needed to test the api.

## Rules

I added the proper logic to satisfy the rules for the receipt processor.

## Requirements

- **Go Installed:** This project is designed to run natively in Go version 1.23.3.
- **Docker Installed:** If not using Go directly, a Dockerized setup is provided for convenience.

## Running the application via Docker

1. Clone the repository and navigate to the project directory.
2. Build the Docker image:
   ```bash
   docker build -t receipt-processor .
   ```
3. Run the Docker container:
   ```bash
   docker run -p 8080:8080 receipt-processor
   ```

**NOTE:** The application is configured to use the Docker standard container endpoint, `8080`. If you are running the application on a different port, you will need to update the Postman collection's environment variables to point to the correct port during testing.

## Testing

I have included a Postman [collection](https://www.postman.com/lunar-module-operator-54317393/fetch-take-home/collection/cemgrz2/fetch-take-home?action=share&creator=23899333) for testing the endpoints. Simply fork the collection and import it into your Postman workspace.

 **NOTE:** There is a Postman script included in the Process Receipt request that will automatically set the Receipt ID upon a successful request.

## Thank You

I wanted to thank Fetch for the opportunity to work on this assessment. I learned a lot about Go and Docker in the process. I hope to hear from you soon.

