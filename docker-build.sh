docker build -t fduyh2021/video:v1 -f VideoDockerfile .
docker build -t fduyh2021/user:v1 -f UserDockerfile .
docker build -t fduyh2021/api:v1 -f ApiDockerfile .
docker run -p 8080:8082 --network douyin fduyh2021/api:v1
docker run -p 8081:8081 --network douyin fduyh2021/video:v1
docker run -p 8889:8889 --network douyin fduyh2021/user:v1