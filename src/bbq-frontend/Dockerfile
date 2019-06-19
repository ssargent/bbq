# stage: 1
FROM node:8 as react-build
WORKDIR /app
COPY . ./
RUN yarn
RUN yarn build

# stage: 2 
FROM nginx:alpine
COPY --from=react-build /app/default.conf /etc/nginx/conf.d/default.conf 
COPY --from=react-build /app/build /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
