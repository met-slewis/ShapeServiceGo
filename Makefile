
# Parameters
REGION=ap-southeast-2
PROFILE=coredev-dev
AWS=aws --profile $(PROFILE) --region $(REGION)
BUCKET_NAME=bucket
FILENAME=locations.json

copy-locations:
	${AWS} s3 cp ./res/${FILENAME} s3://${BUCKET_NAME}/${FILENAME}

