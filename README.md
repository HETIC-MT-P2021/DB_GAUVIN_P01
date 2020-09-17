# DB_GAUVIN_P01
Shop DB in Golang

Import db : docker exec -i db_gauvin_p01_db_1 sh -c 'exec mysql -ugobdd -p"gobdd" image_gobdd' < ./docker/data/database.sql

Lauch : docker-compose up --build