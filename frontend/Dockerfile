FROM node:24-alpine as deps
WORKDIR /app

COPY package.json package-lock.json* yarn.lock* ./
RUN npm install

FROM node:24-alpine as builder
WORKDIR /app

COPY --from=deps /app/node_modules ./node_modules
COPY . .

RUN npm run build

FROM node:24-alpine

COPY --from=builder /app ./

EXPOSE 3000

CMD [ "npm", "start" ]