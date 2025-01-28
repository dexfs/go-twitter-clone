aws ecr get-login-password --region us-east-1 --profile bia | docker login --username AWS --password-stdin 174597536931.dkr.ecr.us-east-1.amazonaws.com
docker build -f ../Dockerfile -t  dexfs/go-twitter-clone:latest ./..
docker tag dexfs/go-twitter-clone:latest 174597536931.dkr.ecr.us-east-1.amazonaws.com/dexfs/go-twitter-clone:latest
docker push 174597536931.dkr.ecr.us-east-1.amazonaws.com/dexfs/go-twitter-clone:latest