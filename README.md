# Money Manager
Expense tracking and personal finance.

## Technologies
- [Go](https://golang.org/)
- [AWS Lambda](https://aws.amazon.com/lambda/)
- [AWS DynamoDB](https://aws.amazon.com/dynamodb/)
- [AWS SAM CLI](https://aws.amazon.com/serverless/sam/)
- [Serverless Framework](https://www.serverless.com/)
- [Golangci-lint](https://golangci-lint.run/)
<br />

## Getting Started
### Requirements
1. [Install](https://golang.org/doc/install) Go > 1.16
2. [Install](https://nodejs.org/en/download/) NodeJS > 14.16
3. [Install](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html) SAM CLI > 1.21
4. [Install](https://www.serverless.com/framework/docs/getting-started/) Serverless Framework CLI > 2.31
5. [Install](https://golangci-lint.run/usage/install/#local-installation) Golangci-lint > 1.31
6. Install Make: > 4.2
    - Linux/Debian: `sudo apt-get install build-essential`
    - Windows: `choco install make`
    - Mac: `brew install make`
7. Run: `make install`
<br />

## Development
### Build
- `make build`

### Build and Run
- `make run`

### Clean cache
- `make clean`

### Golangci-linter
- `make lint`
