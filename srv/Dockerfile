FROM php:7.4-apache
# RUN docker-php-ext-install mysqli && docker-php-ext-enable mysqli
RUN apt-get update && apt-get upgrade -y

# RUN docker-php-ext-install mysqli 
# RUN docker-php-ext-enable mysqli
# RUN apachectl restart

COPY ./php.ini-production  .
RUN mv "php.ini-production" "$PHP_INI_DIR/php.ini"  