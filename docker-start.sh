sudo docker network create uniting_net

sudo docker-compose up -d --build

sudo docker pull testtask
sudo docker pull postgresDB

sudo docker-compose up