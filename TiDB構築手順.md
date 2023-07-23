## TiDB ローカル環境構築 手順

- [ ] 0. 事前準備

  - 0.1 docker imageの参照先変更

    ~~~bash
        eval $(minikube docker-env)
        kubectl config get-contexts
    ~~~
  
  - 0.2  minikubeのメモリを設定する

    ~~~bash
        # Minikubeのメモリを設定(約8GB)
        minikube config set memory 7960
    ~~~
  
  - 0.3 minikubeを再起動する

    ~~~bash
        minikube start
    ~~~

- [ ] 1. 導入要件であるhelm3をインストール

    ~~~bash
        brew install helm
    ~~~

- [ ] 2. TiDBをデプロイ
  
    ~~~bash
        kubectl create** -f https://raw.githubusercontent.com/pingcap/tidb-operator/v1.3.9/manifests/crd.yaml
    ~~~

- [ ] 3. tidb-operatorをhelmでデプロイ

    ~~~bash
        helm repo add pingcap https://charts.pingcap.org/
        kubectl create namespace tidb-admin
        helm install --namespace tidb-admin tidb-operator pingcap/tidb-operator --version v1.3.9
    ~~~

- [ ] 4. デプロイできたかを確認

    ~~~bash
        kubectl get pods --namespace tidb-admin -l app.kubernetes.io/instance=tidb-operator
    ~~~

    **少々待つ**

- [ ] 5. TiDBを作成

    ~~~bash
        kubectl create namespace sample
        kubectl apply --namespace sample -f sample-tidb-cluster.ya
        kubectl get pod --namespace sample
    ~~~

- [ ] 6. TiDBに接続

  - 6.1 ロードバランサのIPを確認

    ~~~bash
        kubectl get service --namespace sample
    ~~~

  - 6.2 TiDBに接続する

    ~~~bash
        mysql --comments -h [ロードバランサのexternal IP] -P 4000 -u root
    ~~~

- [ ] 7. TiDBの削除
  
  ~~~bash
    kubectl delete deployment server デプロイメント名
  ~~~

