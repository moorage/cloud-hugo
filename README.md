# Chugo

A CMS which internally uses hugo and provides a clean and user friendly way to edit, build and deploy static sites

# How to

- Build the binary
```
go build ./cmd/chugo/
```
- we use google pub/sub so create a service account with `Pub/Sub Editor` role and download the json
- rename that file to `credentials.json` and place it in the same dir as the `chugo` binary
- run the chugo binary from the root folder
- download caddy by running the below commands in cloud-hugo dir or run the `./scripts/download-caddy.sh`
```bash
wget -O /tmp/caddy.tar.gz https://caddyserver.com/download/linux/amd64?license=personal
tar -xzf /tmp/caddy.tar.gz --directory /tmp
cp /tmp/caddy ./
```
- now run the `caddy` binary
- now publish a message of kind
```
{
    "git_url": "<git url>"
}
```

eg - https://github.com/cloudacademy/static-website-example.git

- now goto [http://localhost:8081/\[repo-name\]](http://localhost:8081/static-website-example)