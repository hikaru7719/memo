FROM node:10
RUN mkdir /app
WORKDIR /app
COPY package*.json ./
RUN yarn install
CMD yarn dev
