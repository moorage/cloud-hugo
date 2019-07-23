# Chugo

A CMS which internally uses hugo and provides a clean and user friendly way to edit, build and deploy static sites

## Pre-requsites

 - npm 6.4.1
 - node v8.16.0
 - go 1.12.5
 - dep

# Building

Before building anything,we use google pub/sub so create a service account with `Pub/Sub Editor` role and download the from google cloud. Rename that file to `credentials.json` and place it in the same dir as the publisher and subscriber ( both need this credential file ). 

## How to run the publisher
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
    "user_name": "johndoe",
    "user_email": "johndoe@gmail.com",
    "access_token": "<<github token>>",
    "repo_url": " https://github.com/girishramnani/hugo-sample.git"
}
```

## How to run the subscriber

- Build the binary
```
go build ./cmd/subscriber/
```

- run the subscriber binary from the root folder
- download caddy by running the below commands in cloud-hugo dir or run the `./scripts/download-caddy.sh`
```bash
wget -O /tmp/caddy.tar.gz https://caddyserver.com/download/linux/amd64?license=personal
tar -xzf /tmp/caddy.tar.gz --directory /tmp
cp /tmp/caddy ./
```
- now run the `caddy` binary
- open a new console and run the `subscriber` binary in the same directory as `caddy`

eg - https://github.com/cloudacademy/static-website-example.git

- now goto [http://localhost:8081/](http://localhost:8081)

### Configuration for subscriber

- a file called `sub-config.json` has to be provided in the folder same as the subscriber. e.g.
```
{
    "repo_url": " https://github.com/girishramnani/hugo-sample.git"
}
```
