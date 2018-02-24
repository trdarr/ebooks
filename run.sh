git fetch
git reset --hard origin/master
docker build -t ebooks:latest -t ebooks:$(git rev-parse --short HEAD) .
docker stop ebooks
docker rm ebooks
docker run -d --name ebooks --network ebooks ebooks:latest
