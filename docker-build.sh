docker build -t fduyh2021/video:v1 -f VideoDockerfile .
docker build -t fduyh2021/user:v1 -f UserDockerfile .
docker build -t fduyh2021/api:v1 -f ApiDockerfile .
docker run --name video -p 8088:8088 --network douyin fduyh2021/video:v1
docker run --name user -p 8087:8087 --network douyin fduyh2021/user:v1
docker run --name api -p 8089:8089 --network douyin fduyh2021/api:v1