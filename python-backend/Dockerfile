FROM python:3.7-alpine3.12 as build

# Working directory
WORKDIR /root/

# Copy dependencies
COPY requirements.txt .

# Install utils dependencies for Mariadb
RUN apk add build-base \
    && apk add python3-dev \
    && apk add musl-dev \
    && apk add mariadb-dev

# Update pip
RUN python -m pip install --upgrade pip

# Install dependencies
RUN pip install -r requirements.txt

# J'ai download django à part car j'avais une erreur étrange qui doit provenir de mon mac.
RUN pip install --upgrade Django

# Copy rest of the app
COPY . .

EXPOSE 8000

RUN chmod +x ./start.sh

# Exec
CMD [ "./start.sh" ]