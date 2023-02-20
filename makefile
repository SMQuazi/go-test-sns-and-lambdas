export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0
.DEFAULT_GOAL := deploy

deploy: 
	go build main.go
	zip -r main.zip main
	aws lambda update-function-code --function-name "GoLambdaTest" --zip-file fileb://main.zip --region="us-east-2"