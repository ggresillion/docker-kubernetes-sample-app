FROM node:alpine as builder

RUN npm i -g pnpm

WORKDIR /app

COPY package.json pnpm-lock.yaml /app/

RUN pnpm i

COPY . .

RUN pnpm run build

FROM node:alpine

WORKDIR /app

COPY --from=builder /app/build /app/package.json /app/

EXPOSE 3000

CMD ["node", "/app"]
