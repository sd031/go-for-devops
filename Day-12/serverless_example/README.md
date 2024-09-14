# Run below if you are doing from scratch
``` bash
go mod init serverless_example
go get github.com/aws/aws-lambda-go/lambda
go run main.go
```


# Run  if already set-up

```bash
go mod tidy
go run main.go
```

# Test Locally 
1) Install AWS Serverless Application Model Command Line Interface (AWS SAM CLI) : 
https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/install-sam-cli.html

2) Create template.yaml for SAM: Define your Lambda function in a template.yaml file:
```yaml
AWSTemplateFormatVersion: '2010-09-09'
Transform: AWS::Serverless-2016-10-31
Resources:
  GoLambdaFunction:
    Type: AWS::Serverless::Function
    Properties:
      Handler: main
      Runtime: go1.x
      CodeUri: .
      MemorySize: 128
      Timeout: 10
      Environment:
        Variables:
          PARAM_NAME: VALUE
```
3) Build the Lambda function to test locally
```bash
sam build
```

4) Run the Lambda function locally
Login in Public ECR:
```bash
aws ecr-public get-login-password --region us-east-1 | docker login --username AWS --password-stdin public.ecr.aws
```
```bash
sam local invoke GoLambdaFunction --event event.json
```
if mac:
```bash
DOCKER_HOST=unix://$HOME/.docker/run/docker.sock sam local invoke GoLambdaFunction --event event.json
```




5) Deploy uisng SAM
```bash
sam deploy --guided
```

6) Test remotly / cloud after deploy:
```bash
sam remote invoke GoLambdaFunction --stack-name sam-app --event '{"name": "Sandip"}'
```


# Build the Lambda function & Deploy Directly :
1) Build the binary
```bash
GOOS=linux GOARCH=amd64 go build -o main main.go
```

2) Zip the binary:
```bash
zip function.zip main
```

3) Deploy to AWS Lambda using AWS CLI:
```bash
# Create IAM role
aws iam create-role --role-name LambdaExecutionRole --assume-role-policy-document '{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}'

# Attach policy
aws iam attach-role-policy --role-name LambdaExecutionRole --policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole

# Get Role ARN
ROLE_ARN=$(aws iam get-role --role-name LambdaExecutionRole --query 'Role.Arn' --output text)

# Create Lambda function
aws lambda create-function --function-name goLambdaFunction --zip-file fileb://function.zip --handler main --runtime go1.x --role $ROLE_ARN

```

4) Test by invoking the function 
```bash
aws lambda invoke --function-name goLambdaFunction --payload '{"name": "Sandip"}' output.json
cat output.json
```




