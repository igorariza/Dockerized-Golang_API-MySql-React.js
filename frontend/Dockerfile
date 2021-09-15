FROM node:alpine


WORKDIR /frontend


COPY ["package.json", "package-lock.json*", "./"]



RUN npm i



EXPOSE 3000

CMD ["npm", "run", "start"]