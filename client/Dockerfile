FROM node:18.13.0

COPY . ./var/hause/client
WORKDIR /var/hause/client

EXPOSE 5173

RUN npm install
CMD [ "npm start" ]
