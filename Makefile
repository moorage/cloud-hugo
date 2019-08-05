

TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)
export TAG

build_front:
	cd frontend
	npm install
	npm run build

build_pub: build_front
	go build ./cmd/publisher

build_sub:
	go build ./cmd/subscriber

createCluster: 
	gcloud config set project ${GOOGLE_PROJECTID}
	gcloud container clusters create ${GOOGLE_PROJECTID}-cluster  --zone "us-central1-a" --no-enable-basic-auth --cluster-version "1.12.8-gke.10" --machine-type "n1-standard-1"


build: build_pub build_sub

pack: build_pub build_sub
	docker build -f Dockerfile.pub -t gcr.io/clugo/publisher:$(TAG) .
	docker build -f Dockerfile.sub -t gcr.io/clugo/subscriber:$(TAG) .

upload: pack
	docker push gcr.io/clugo/publisher:$(TAG)
	docker push gcr.io/clugo/subscriber:$(TAG)