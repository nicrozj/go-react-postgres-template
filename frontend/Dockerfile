FROM node:20-alpine AS dev
WORKDIR /app
COPY package*.json ./
RUN npm install
CMD ["npm", "run", "dev"]