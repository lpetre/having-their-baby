#!/bin/bash -xe

TMPDIR=$(mktemp -d -t lambda-XXXXXXXXXX)
trap '{ rm -rf $TMPDIR; }' EXIT

cp credentials.json "$TMPDIR"
GOOS=linux go build -o "$TMPDIR/main" main.go

cd "$TMPDIR"
zip function.zip main credentials.json
aws lambda update-function-code --region us-east-1 --function-name  arn:aws:lambda:us-east-1:039145332735:function:twilioReceiveHavingTheirBaby --zip-file "fileb://$TMPDIR/function.zip"