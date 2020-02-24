sudo docker network create uniting_net

sudo docker-compose up -d --build

sudo docker pull test_task
sudo docker pull postgres_db

cd db

sql-migrate up

sudo docker-compose up
