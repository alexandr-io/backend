docker build -t user .
docker run -p 3000:3000 --rm -d --name user user:latest