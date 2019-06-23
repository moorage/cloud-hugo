# Chugo

A CMS which internally uses hugo and provides a clean and user friendly way to edit, build and deploy static sites

# How to

- Build the binary
```
go build ./cmd/chugo/
```
- we use google pub/sub so create a service account with `Pub/Sub Editor` role and download the json
- rename that file to `credentials.json` and place it in the same dir as the `chugo` binary
