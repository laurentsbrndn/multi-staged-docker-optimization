# FROM composer:2.6 AS vendor

# WORKDIR /app

# COPY composer.json composer.lock ./
# RUN composer install --no-dev --optimize-autoloader --prefer-dist

# FROM php:8.2-fpm-alpine AS app

# RUN apk add --no-cache \
#     libpng-dev \
#     libjpeg-turbo-dev \
#     freetype-dev \
#     oniguruma-dev \
#     libxml2-dev \
#     zip \
#     unzip \
#     icu-dev \
#     curl \
#     && docker-php-ext-install \
#         pdo \
#         pdo_mysql \
#         mbstring \
#         bcmath \
#         gd \
#         intl

# COPY --from=composer:2.6 /usr/bin/composer /usr/bin/composer

# WORKDIR /var/www

# COPY . .

# RUN composer install --no-dev --optimize-autoloader --prefer-dist

# RUN chown -R www-data:www-data /var/www \
#     && chmod -R 775 storage bootstrap/cache \
#     && php artisan config:cache \
#     && php artisan route:cache \
#     && php artisan view:cache

# EXPOSE 8080

# CMD ["php", "artisan", "serve", "--host=0.0.0.0", "--port=8080"]

FROM composer:2.6 AS vendor

WORKDIR /app

COPY composer.json composer.lock ./
RUN composer install --no-dev --optimize-autoloader --prefer-dist --no-scripts

FROM php:8.2-fpm-alpine AS app

RUN apk add --no-cache \
    libpng-dev \
    libjpeg-turbo-dev \
    freetype-dev \
    oniguruma-dev \
    libxml2-dev \
    zip \
    unzip \
    icu-dev \
    curl \
    && docker-php-ext-install \
        pdo \
        pdo_mysql \
        mbstring \
        bcmath \
        gd \
        intl

COPY --from=composer:2.6 /usr/bin/composer /usr/bin/composer

WORKDIR /var/www

COPY --from=vendor /app/vendor ./vendor
COPY . .

RUN composer run-script post-autoload-dump || true

RUN chown -R www-data:www-data /var/www \
    && chmod -R 775 storage bootstrap/cache

RUN php artisan config:cache \
    && php artisan route:cache \
    && php artisan view:cache

EXPOSE 8080

CMD ["php", "artisan", "serve", "--host=0.0.0.0", "--port=8080"]