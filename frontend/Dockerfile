FROM node:lts

WORKDIR /app
COPY . /app/
COPY public /app/public
RUN npm install --frozen-lockfile --production=false
EXPOSE 3000

CMD [ "npm", "start" ]
