[Leia isto em Português](./README-pt-br.md)

# Liqi Backend Challenge Project

This repository contains a solution to Liqi's Backend Challenge. The task is designed to test proficiency in backend development, familiarity with AWS services, and understanding of key concepts in cryptography and blockchain technology.

## Overview
The final project is essentially a backend service that interacts with the Ethereum Network to sign transactions. It represents my first or second experience with Go language and the process of signing transactions on the Blockchain. Given that my primary background is in Node.js, transitioning to Go was an exciting and challenging new experience.

## Learning Journey
The main challenges I faced were adapting to the syntax of Go and researching suitable packages for this project. Below, I've listed some of the resources I used in my learning process:

1. **Go Structs:** One of the fundamental concepts in Go, which helped me understand the structure and types in the language. [Go by Example](https://gobyexample.com/structs) was an excellent resource for this.
2. **JSON operations in Go:** The [encoding/json](https://pkg.go.dev/encoding/json#Unmarshal) package in Go was instrumental in handling JSON operations.
3. **ECDSA (Elliptic Curve Digital Signature Algorithm):** The [crypto/ecdsa](https://pkg.go.dev/crypto/ecdsa) package in Go helped me understand the ECDSA used for the cryptographic requirements of the challenge.

For the Ethereum node, I utilized services from [Infura](https://infura.io/), and AWS services (SQS and Lambda) for handling messages and executing functions respectively. The [Serverless Framework](https://www.serverless.com/) proved to be a valuable tool for deployment.

## Project Stages

The challenge was split into two parts:

1. **REST API Development:** I developed a REST API with two endpoints: one for generating private-public keys using ECDSA, and another for obtaining an Ethereum EOA (Externally Owned Account) address from a public key. All functionalities were covered with unit tests.

2. **AWS Lambda and SQS Setup:** In the second part, I created an AWS Lambda function triggered every time a new message arrived in an AWS SQS queue. The SQS messages are Ethereum transactions which the Lambda function signs with a generated private key and sends to an EVM (Ethereum Virtual Machine) blockchain.

The challenge took me approximately 20 hours spread over three business days to complete. While the code may not be perfect due to my relative newness to Go, I believe it successfully meets the requirements of the task and demonstrates my learning journey.

## Running Locally

To run the project locally using Docker, follow the steps below:

1. **Clone the Repository**
    ```
    git clone git@github.com:philipep-galdino/liqi-challenge-project.git
    ```

2. **Navigate to Project Directory**
    ```
    cd liqi-challenge-project
    ```

3. **Build the Docker Image**
    ```
    docker build -t liqi-challenge-project .
    ```

4. **Run the Docker Container**
    ```
    docker run -p 8080:8080 liqi-challenge-project
    ```

Once the application is running, you can use an API testing tool like Postman or cURL to make requests to your local server at `http://localhost:8080`.

## Deployment

Before starting, ensure that the Serverless Framework is installed and configured on your system. If not, you can follow the guide on their official documentation: https://www.serverless.com/framework/docs/getting-started/

You also need to set up your own AWS SQS queue and update the `serverless.yml` file with your specific AWS SQS queue details.

To deploy the application, follow the steps below:

1. **Set Up Your Environment Variables**
   
    Create a `.env` file in the root directory and set the necessary environment variables. For instance:

    ```
    AWS_ACCESS_KEY_ID=<Your AWS Access Key>
    AWS_SECRET_ACCESS_KEY=<Your AWS Secret Access Key>
    SQS_QUEUE_URL=<Your AWS SQS Queue URL>
    ```
   
2. **Deploy with Serverless**
   
    In the root directory of your project, run:

    ```
    serverless deploy
    ```

This will package and deploy your application to AWS Lambda. Once deployment is successful, Serverless Framework will provide you with the deployed API endpoints.


