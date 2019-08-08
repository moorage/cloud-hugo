# Chugo

A CMS which internally uses hugo and provides a clean and user friendly way to edit, build and deploy static sites

## Pre-requsites

 - npm 6.4.1
 - node v8.16.0
 - go 1.12.5
 - dep

## Cloning

As we have git submodules use the below to clone the sub modules as well
```
git clone --recurse-submodules -j8 git@github.com:moorage/cloud-hugo.git
```

# Building

Before building anything,we use google pub/sub so create a service account with `Pub/Sub Editor` role and download the from google cloud. Rename that file to `credentials.json` and place it in the same dir as the publisher and subscriber ( both need this credential file ). 

## How to run the publisher

- Ensure you're dependences are up to date
```
dep ensure
```
- Build the binary
```
go build ./cmd/publisher
```
- build the stackedit editor based frontend
```
cd frontend
npm install
npm run build
cd ..
```
- run the publisher binary
- go to `http://locahost:8080` which will open the stack edit editor for you

### Configuration for publisher

- a file called `pub-config.json` has to be provided in the folder same as the publisher. e.g.
```
{
    "project_id": "cloud-hugo-test",
    "topic_name": "chugo-run-requests",
    "env": "dev",
    "pub_port": "8080",
    "user_name": "<<username>>",
    "user_email": "<<email>>",
    "access_token": "<<token>>",
    "repo_url": " https://github.com/girishramnani/hugo-sample.git"
}
```

## How to run the subscriber

- Build the binary
```
go build ./cmd/subscriber/
```

- run the subscriber binary from the root folder
- now goto [http://localhost:8081/](http://localhost:8081)

### Configuration for subscriber

- a file called `config/sub-config.json` has to be provided in the folder same as the subscriber. e.g.
```
{
    "project_id": "cloud-hugo-test",
    "topic_name": "chugo-run-requests",
    "env": "dev",
    "sub_port": "8080",
    "user_name": "<<username>>",
    "user_email": "<<email>>",
    "access_token": "<<token>>",
    "repo_url": " https://github.com/girishramnani/hugo-sample.git"
}
```

## Deploying to kubernetes

The files related to kubernetes are present in `k8s/` - 

- `config.yml` - This file holds the pub-config and sub-config we populate with environment variables
- `pub-deployment.yml` - the publisher deployment config
- `sub-autoscaling-yml` - the autoscaling deployment for subscriber

to build and push and run all the stuff on kuberenetes -

1 - set all env variable present in the `Makefile` 
```
export GOOGLE_PROJECTID?=cloud-hugo-test
export GITHUB_USERNAME?=<<username>>
export GITHUB_ACCESS_TOKEN?=<<token>>
export GITHUB_EMAIL?=<<email>>
export REPO_URL?=<<repo_url>>
```

2 - add the `credentials.json` (this is the google cloud creds for pub sub) in the `config/` folder 

3 - `make deploy`

## build the frontend

Currently we ship a prebuild frontend which is part of the repo as a submodule in the folder "frontend".
In case you ever want build the frontend because you have changed something the you can do so by
```
make buildFront
```