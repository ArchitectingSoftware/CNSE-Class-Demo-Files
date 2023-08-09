#! /bin/bash
#CREATES A KUBERNETES CLUSTER CALLED WP-VUL
kind create cluster --name cnse-class --config=kind-config.yml

kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml
kubectl wait --namespace ingress-nginx \
  --for=condition=ready pod \
  --selector=app.kubernetes.io/component=controller \
  --timeout=90s