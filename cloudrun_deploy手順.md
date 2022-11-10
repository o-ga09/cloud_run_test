## Cloud Run deploy 手順


 シンプルなGo Web サーバをGoogle Cloud Run にデプロイする時の手順


- [ ] 0. 事前準備
    不要なDockerイメージの削除など
~~~bash
docker image ls
docker image rmi gcr.io/cloud-run-test 
docker image rmi taiti09/cloud-run-test
docker container rm taiti09/cloud-run-test
docker container rm 5c83bb584415
docker image rmi taiti09/cloud-run-test
~~~

- [ ] 1. デプロイ用のDockerイメージをビルド

~~~bash
docker build -t gcr.io/todo-app-20221107/
cloud-run-test:latest --target deploy ./
~~~

- [ ] 2. GCPのCloud Registry にプッシュ

~~~bash
gcloud builds submit --tag gcr.io/todo-app-20221107/cloud-run-test
~~~

- [ ] 3. Cloud Run にデプロイ
~~~bash
gcloud run deploy --image gcr.io/todo-app-20221107/cloud-run-test --platform managed --region us-west1 
~~~

- [ ] 4. Cloud Runのサービスを削除

~~~bash
gcloud run services delete --platform managed --region us-west1 simple-go-web-server
~~~