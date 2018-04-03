default:
	GOOS=linux GOARCH=amd64 go build -o main main.go && zip main.zip main

package:
	aws cloudformation package --template-file template.yaml --output-template-file=deploy.yaml --s3-bucket $(S3_BUCKET_NAME)

deploy:
	aws cloudformation deploy --template-file deploy.yaml --stack-name $(CF_STACK_NAME) --capabilities CAPABILITY_IAM