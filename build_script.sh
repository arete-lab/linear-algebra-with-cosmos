docker stop linear-algebra && docker rm linear-algebra
docker build -f Dockerfile . -t linear-algebra
docker create --name linear-algebra -i -v $(pwd):/linear-algebra -w /linear-algebra -p 1317:1317 -p 3000:3000 -p 4500:4500 -p 5000:5000 -p 26657:26657 linear-algebra
docker start linear-algebra
docker exec -it linear-algebra ignite chain serve