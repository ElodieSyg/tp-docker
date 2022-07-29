FROM node:18.7.0-alpine3.16 as developpement

# Working directory
WORKDIR /root/

# Copy dependencies
COPY package*.json .

# Install dependencies
RUN npm install

# Copy rest of the app
COPY . .

# Run build for production
RUN npm run build

# Import nginx image
FROM nginx:1.23.1-alpine as production

# Copy static html file
COPY --from=developpement /root/build /usr/share/nginx/html

EXPOSE 80

ENTRYPOINT ["nginx", "-g", "daemon off;"]
# Production
# COMD ["npm", "ci", "as=production"]