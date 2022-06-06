docker build -t fduyh2021/video:v1 -f VideoDockerfile .
docker build -t fduyh2021/user:v1 -f UserDockerfile .
docker build -t fduyh2021/api:v1 -f ApiDockerfile .
docker run -p 8087:8087 --network douyin fduyh2021/api:v1
docker run -p 8088:8088 --network douyin fduyh2021/video:v1
docker run -p 8089:8089 --network douyin fduyh2021/user:v1