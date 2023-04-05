#使うDockerイメージ
FROM mysql

#ポートを開ける
EXPOSE 8080

#MySQL設定ファイルをイメージ内にコピー
ADD ./docker/db/my.cnf /etc/mysql/conf.d/my.cnf

#docker runに実行される
CMD ["mysqld"]