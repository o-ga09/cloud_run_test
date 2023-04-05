## k8s ローカル環境構築 手順


 シンプルなGo Web サーバをローカルのk8s にデプロイする時の手順


- [ ] 0. 事前準備
    - 0.1 minikube、kubectlのインストール
    ~~~bash
        brew install minikube
        rew install kubectl
    ~~~

    - 0.2 docker imageの参照先変更

    ~~~bash
        eval $(minikube docker-env)
        kubectl config get-contexts
    ~~~

    - 0.3 docker compose からイメージ削除

    ~~~bash
        docker-compose down --rmi all --volumes --remove-orphans
    ~~~

- [ ] 1. start コマンドで単一ノードの kubernetes クラスタを作成

~~~bash
minikube start
~~~

- [ ] 2. クラスタの稼働状況を確認
    status コマンドで minikube の稼働状況を見ることができます。
    Running になっていれば起動に成功しています。

~~~bash
minikube status
~~~

- [ ] 3. クラスタの停止、削除

~~~bash
minikube stop
minikube delete
~~~

- [ ] 4.  kubectlによるクラスタの操作
    - 4.1 ノードの取得

    ~~~bash
        kubectl get node
    ~~~
    - 4.2 ノードの情報を取得
        cluster-info コマンドでクラスタの情報を確認することができます。コントロールプレーンのドメイン情報などが表示されます。
    ~~~bash
        kubectl cluster-info 
    ~~~

    - 4.3 ポッドの起動
        ポッドをコントローラーを通さずに起動してみます。
        -it を指定しないとバッググラウンドで起動
    ~~~bash
        kubectl run ポッド名 --image=コンテナイメージ名 -it --restart=Never
    ~~~

    
    - 4.3.1
        ポッドの情報取得

        ~~~bash
            kubectl describe pod env-server
        ~~~

    - 4.4 ポッドの取得
    
    get の後ろに deploy,pod 等項目を指定できる
    ~~~bash
        kubectl get pod
    ~~~

    - 4.5 ポッドの削除

    ~~~bash
        kubectl delete pod コンテナイメージ名
    ~~~

    - 4.6 ポッドのログを表示

    ~~~bash
        kubectl logs コンテナイメージ名
    ~~~

- [ ] 5. デプロイメントを通してポッドを起動

    - 5.1 デプロイメントを作成

        create deployment コマンドでデプロイメントを作成できます。これにより、デプロイメントを通してポッドを起動する事ができます。
        1. デプロイメントを作成する
        2. デプロイメントがポッドの目標数を維持しようとする
        3. 目標数が1であれば、ポッドを1個作成する
        
        ~~~bash
        kubectl create deployment --image=コンテナイメージ名 ポッド名 デプロイメント名/ポッド名 created
        ~~~
    
    - 5.2 ポッドのレプリカを変更する

    ~~~bash
        kubectl scale --replicas=5 deployment/server
    ~~~

    - 5.3 ポッドの削除

    ~~~bash
        kubectl delete pod ポッド名１ ポッド名２
    ~~~

    - 5.5 デプロイメントの削除

    ~~~bash
        kubectl delete deployment server デプロイメント名
    ~~~

    - 5.5 デプロイメント定義を作成する

    ~~~bash
        kubectl create -f ymalファイル名
    ~~~

    - 5.6 デプロイメント定義を適用する

    ~~~bash
        kubectl apply -f ymalファイル名
    ~~~

- [ ] 6.サービスの起動

    ~~~bash
        minikube service go-app-service
    ~~~

- [ ] 7. Ingressでサービスを外部公開する

    minikubeでIngressを有効にする

    ~~~bash
        minikube addons enable ingress
        minikube tunnel
    ~~~

    IngressのPodが起動していることを確認する

    ~~~bash
        kubectl get pods -n kube-system | grep ingress
    ~~~

## 注意
  - yamlファイルは半角英数字と-のみ　(_ は使えない)
  - Go でlog.Fatalf は気軽使わない。(内部でos.exit(1)をしている)
  - base64 エンコード

    ~~~bash
        echo -n "giraffi" | base64  
    ~~~
  - base64 デコード

    ~~~bash
        echo -n "giraffi" | base64 -d
    ~~~