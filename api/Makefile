.DEFAULT_GOAL := run
.PHONY : mongo  benchmark run  
mongo : 
	docker run -d --name mongodb -p 27017:27017 -v mongodb_data:/data/db \
    -e MONGO_INITDB_ROOT_USERNAME=myuser \
    -e MONGO_INITDB_ROOT_PASSWORD=mypassword \
    mongo
benchmark :
    ab -n 5000 -c 200 -g with-cache.data http://localhost:8081/api/v1/recipes

run : 
    JWT_SECRET=B1OZ2f3pQJzGS2ctvmt5zQ== MONGO_URI=mongodb://myuser:mypassword@localhost:27017/ MONGO_DATABASE=demo go run cmd/*.go   

