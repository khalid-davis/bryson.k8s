all: kubernetes-tester
#0.1.0

TAG = 0.1
PREFIX = kevindavis/kubernetes-tester

vet:
	go vet ./...

fmt:
	go fmt ./...

kubernetes-tester: *.go controllers/*.go routers/*.go models/*.go k8s/*.go vet fmt
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o $@

imagesave:
	docker build -t $(PREFIX):$(TAG) .
	docker save $(PREFIX):$(TAG) >kubernetes-tester-$(TAG).tar
	gzip kubernetes-tester-$(TAG).tar

push:
	docker build -t $(PREFIX):$(TAG) .
	docker push $(PREFIX):$(TAG)

clean:
	rm -f kubernetes-tester
	rm -f kubernetes-tester-$(TAG).tar.gz
