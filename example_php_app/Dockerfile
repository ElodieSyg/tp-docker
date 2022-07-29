FROM php:7.2-apache as build

# Working directory
WORKDIR /root

# Install git to make composer works
RUN apk --update --no-cash add git

# Copy dependencies
COPY composer.json .
COPY composer.lock .


# COPY /public/.htaccess .

# Dowload dependencies
RUN composer update
RUN composer install

# Copy rest of the app
COPY . .

FROM composer:2.3.10



EXPOSE 7869

CMD [ "executable" ]

# multistage composer
# copy conf php in public
# lancer apache