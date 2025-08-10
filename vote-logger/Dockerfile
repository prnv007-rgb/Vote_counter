FROM ruby:3.2-alpine

WORKDIR /app
COPY . .

# Install required build tools for native extensions
RUN apk add --no-cache build-base

# Install Sinatra and required gems directly
RUN gem install sinatra rackup puma

EXPOSE 4567
CMD ["ruby", "logger.rb", "-o", "0.0.0.0"]