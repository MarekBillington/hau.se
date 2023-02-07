FROM node:18.13.0-alpine AS build

WORKDIR /app
COPY package*.json ./
RUN npm install
COPY ./ .
RUN npm run build

FROM node:18.13.0-alpine AS prod

WORKDIR /app

COPY --from=build /app/node_modules ./node_modules
COPY --from=build /app/server ./server
COPY --from=build /app/dist ./dist

EXPOSE 3000

CMD [ "node", "server/entry.express" ]
