# Serverless Template for Golang

This repository contains template for creating serverless services written in Golang.

## Quick Start

1. Create a new service based on this template

```
serverless create -u https://github.com/katsut/awslogs_typetalk/ -p awslogs_typetalk
```

2. Compile function

```
cd awslogs_typetalk
GOOS=linux go build -o bin/main
```

3. Edit environment variables

edit environment block in serverless.yml

3. Deploy

```
serverless deploy
```
