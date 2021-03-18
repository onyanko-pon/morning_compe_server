# Morning Compe のサーバーサイド 

## ローカル開発環境

初期起動
```bash
$ docker-compose build
# $ docker volume create pgsql-data
```

起動
```bash
$ docker-compose up
```

### Goサーバー

```bash
$ docker exec -it sagyo_go_server /bin/bash
```

### postgreSQL

```bash
$ psql -h localhost -p 5432 -U admin -d my_db
```
