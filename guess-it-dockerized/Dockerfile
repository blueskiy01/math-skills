FROM node:21-bookworm-slim

WORKDIR /app/

COPY package.json package-lock.json ./ 

RUN npm install

RUN apt-get update -y

RUN apt-get install golang python3 curl -y

COPY . /app/

RUN cp -r ./student/ ../student/

RUN cp -r ./student/ ~/student/

CMD ["npm", "start"]
