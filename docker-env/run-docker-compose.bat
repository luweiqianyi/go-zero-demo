cd ..
copy /Y .\cmd\account\etc\account-api.yaml .\docker-env\account\etc\
copy /Y .\cmd\account-rpc\etc\account.yaml .\docker-env\account-rpc\etc\
copy /Y .\cmd\account-rpc\etc\redis.yaml .\docker-env\account-rpc\etc\

cd docker-env
docker-compose -p go-zero-demo up -d