# バックエンドフォルダ

## 環境構築
```
cd backend
./Genkey.bat
```

server.key と server.crt を nginx/keys に移動

docker
```
docker compose up -d --build
```

https://localhost:8447/statics/ を開く