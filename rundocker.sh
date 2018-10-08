docker container stop tradingdb
docker container rm tradingdb
docker run -d -p 7888:7888 --name tradingdb -v $PWD/dat:/home/tradingdb/dat tradingdb