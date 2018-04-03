# A Golang based Lambda app

This has a sample Gorilla app deployed with AWS Lambda

## 🔨 Building

To build the lamdba app, use

    make

This will build the main and main.zip files. To then package the app, run the following. It uses the `S3_BUCKET_NAME` environment variable to upload `main.zip` to the specified bucket.

    make package

This will package the app as cloudformation template and write `deploy.yaml` to the local folder. To deploy it run the following. This also requires that you set the `CF_STACK_NAME` environment variable to set the stack name for your deployment.

    make deploy

This will read `deploy.yaml` and run setup the function, IAM and API gateway. Ensure that you have the AWS credentials and region setup.

## 🚀 Running

To run the app, visit the API gateway section on your AWS Console and look at the Prod/Stage URLs. 