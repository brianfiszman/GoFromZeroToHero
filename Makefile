.PHONY: run-debug
run-debug: 
	env `cat .env` go run ./cmd/main.go

.PHONY: create-topics
create-topics:
	kubectl exec kafka-0 -- /opt/bitnami/kafka/bin/kafka-topics.sh --create --replication-factor 1 --partitions 1 --topic users --zookeeper kafka-zookeeper --if-not-exists


.PHONY: skaffold-dev
skaffold-dev:
	skaffold dev --auto-build --auto-deploy --tail --cleanup

.PHONY: skaffold-debug
skaffold-debug:
	skaffold debug --auto-build --auto-deploy --tail --cleanup

.PHONY: encrypt-secrets
encrypt-secrets:
	helm secrets enc k8s/secrets.yaml
