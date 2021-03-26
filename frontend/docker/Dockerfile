FROM node:15.12-alpine

WORKDIR /app

ENV PATH /app/node_modules/.bin:$PATH

COPY package.json /app/package.json
RUN npm install

CMD ["npm", "run", "serve"]