# Next-Go-Template
以下の構成をkubernetesクラスタで構築するためのマニフェスト
- Frontend: Next.js x Typescript  
- Backend: Golang(gin), ORM/gorm  
- DB: postgresql  

## 利用方法
### ローカル環境での実行
minikube, docker-desktop for mac のいずれかをインストール

```
kubectl apply -f infrastructure/db
```

```
kubectl apply -f infrastructure/backend
```

```
kubectl apply -f infrastructure/frontend
```

### その他
docker-desktop for macでは`ingress-controller`を別途インストールする必要あり
```
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.8.2/deploy/static/provider/cloud/deploy.yaml
```
参照：https://kubernetes.github.io/ingress-nginx/deploy/#docker-for-mac
