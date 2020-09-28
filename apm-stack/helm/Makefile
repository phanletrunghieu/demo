.PHONY: generate-keys install

generate-keys:
	docker run --rm -it  -v $(shell pwd)/keys:/keys docker.elastic.co/elasticsearch/elasticsearch:7.8.1 sh -c 'elasticsearch-certutil cert --keep-ca-key --pem --in /keys/instance.yml --out /keys/certs.zip && unzip /keys/certs.zip -d /keys && rm /keys/certs.zip'

install:
	helm upgrade --install apm -n app -f values.yaml .