

export TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)
export GOOGLE_PROJECTID?=cloud-hugo-test
export GITHUB_USERNAME?=girishramnani
export GITHUB_ACCESS_TOKEN?=token
export GITHUB_EMAIL?=girishramnani95@gmail.com
export REPO_URL?=https://github.com/girishramnani/hugo-sample.git

createCluster: 
	gcloud config set project ${GOOGLE_PROJECTID}
	gcloud container clusters create ${GOOGLE_PROJECTID}-cluster  --zone "us-central1-a" --no-enable-basic-auth --cluster-version "1.12.8-gke.10" --machine-type "n1-standard-1"

buildFront:
	cd frontend && npm install && npm run build

buildPub:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"'  ./cmd/publisher

buildSub:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"'  ./cmd/subscriber

build: buildPub buildSub

dockerSub: 
	docker build -f Dockerfile.sub -t gcr.io/${GOOGLE_PROJECTID}/subscriber:$(TAG) .

dockerPub:
	docker build -f Dockerfile.pub -t gcr.io/${GOOGLE_PROJECTID}/publisher:$(TAG) .

docker: dockerSub dockerPub

packSub: buildSub dockerSub

packPub: buildPub dockerPub

pack: build docker

uploadSub: packSub
	docker push gcr.io/${GOOGLE_PROJECTID}/subscriber:$(TAG)

uploadPub: packPub
	docker push gcr.io/${GOOGLE_PROJECTID}/publisher:$(TAG)

upload: uploadPub uploadSub

deploySecret:
	kubectl create configmap clugo-secret --from-file=config/credentials.json

deployConfig:
	envsubst < k8s/config.yml | kubectl apply -f -

deployPub: uploadPub
	envsubst < k8s/pub-deployment.yml | kubectl apply -f -

deploySub: uploadSub
	envsubst < k8s/sub-autoscaling.yml | kubectl apply -f - 

deploy: deploySecret deployConfig deployPub deploySub