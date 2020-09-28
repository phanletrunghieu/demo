.PHONY: increase-load

start:
	minikube start --vm-driver=kvm2 --cpus=2 --disk-size=10GB

enable-metric-minikube:
	minikube addons enable metrics-server

increase-load:
	kubectl run --generator=run-pod/v1 -it --rm load-generator --image=busybox /bin/sh
	# while true; do wget -q -O- http://php-apache.default.svc.cluster.local & done