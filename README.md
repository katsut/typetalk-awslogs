# Send AWS CloudWatch Logs to Typetalk

This repository contains template for creating serverless services written in Golang.

## Quick Start

1. Create a new service based on this template

```
sls create -u https://github.com/katsut/typetalk-awslogs/ -p typetalk_awslogs
```

2. Compile function

```
cd typetalk_awslogs

dep ensure

GOOS=linux go build -o bin/main
```

3. Edit environment variables

edit environment block in serverless.yml

3. Deploy

```
sls deploy
```
