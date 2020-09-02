FROM node:14.9.0
RUN npm install -g @angular/cli
RUN mkdir frontend-app
WORKDIR frontend-app
ADD . .
RUN npm install
EXPOSE 4200
ENTRYPOINT ["ng", "serve"]
