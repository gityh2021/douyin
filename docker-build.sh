docker build -t fduyh2021/video:v2 -f VideoDockerfile .
docker build -t fduyh2021/user:v2 -f UserDockerfile .
docker build -t fduyh2021/api:v2 -f ApiDockerfile .
docker run --name video -p 8088:8088 --network douyin fduyh2021/video:v2
docker run --name user -p 8087:8087 --network douyin fduyh2021/user:v2
docker run --name api -p 8082:8082 --network douyin fduyh2021/api:v1