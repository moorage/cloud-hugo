

export TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)
export GOOGLE_PROJECTID?="cloud-hugo-test"
export GITHUB_USERNAME?="girishramnani"

buildFront:
	cd frontend && npm install && npm run build

buildPub: buildFront
	go build ./cmd/publisher

buildSub:
	go build ./cmd/subscriber

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