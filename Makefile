.PHONY: delete create build

delete:
	kubectl delete -f deploy/sa.yaml
	kubectl delete -f deploy/rbac.yaml
	kubectl delete -f deploy/crd.yaml
	kubectl delete -f deploy/operator.yaml
	kubectl delete service/poppins-operator
	kubectl delete -f deploy/cr.yaml
create:
	kubectl create -f deploy/sa.yaml
	kubectl create -f deploy/rbac.yaml
	kubectl create -f deploy/crd.yaml
	kubectl create -f deploy/operator.yaml
build:
	operator-sdk build schorzz/poppins-operator:develop
	docker push schorzz/poppins-operator

create-ns:
	kubectl create -f deploy/cr.yaml -n poppins
	kubectl create -f deploy/busybox-deployment.yaml -n poppins
delete-ns:
	kubectl delete -f deploy/cr.yaml -n poppins
	kubectl delete -f deploy/busybox-deployment.yaml -n poppins
