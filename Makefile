

export TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)
export GOOGLE_PROJECTID?="cloud-hugo-test"
export GITHUB_USERNAME?="girishramnani"
export GITHUB_ACCESS_TOKEN?="token"
export GITHUB_EMAIL?="girishramnani95@gmail.com"
export REPO_URL?="https://github.com/girishramnani/hugo-sample.git"

buildFront:
	cd frontend && npm install && npm run build

buildPub: buildFront
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"'  ./cmd/publisher

buildSub:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"'  ./cmd/subscriber

createCluster: 
	gcloud config set project ${GOOGLE_PROJECTID}
	gcloud container clusters create ${GOOGLE_PROJECTID}-cluster  --zone "us-central1-a" --no-enable-basic-auth --cluster-version "1.12.8-gke.10" --machine-type "n1-standard-1"


build: buildPub buildSub

docker: 
	docker build -f Dockerfile.pub -t gcr.io/${GOOGLE_PROJECTID}/publisher:$(TAG) .
	docker build -f Dockerfile.sub -t gcr.io/${GOOGLE_PROJECTID}/subscriber:$(TAG) .

pack: build docker

upload: pack
	docker push gcr.io/${GOOGLE_PROJECTID}/publisher:$(TAG)
	docker push gcr.io/${GOOGLE_PROJECTID}/subscriber:$(TAG)

deploySecret:
	kubectl create configmap clugo-secret --from-file=config/credentials.json

deployConfig:
	envsubst < k8s/config.yml | kubectl apply -f -

deployPub:
	envsubst < k8s/pub-deployment.yml

deploySub:
	envsubst < k8s/sub-autoscaling.yml | kubectl apply -f - 