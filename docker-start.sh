sudo docker network create uniting_net

sudo docker-compose up -d --build

sudo docker pull test_task
sudo docker pull postgres_db

sudo docker-compose up
