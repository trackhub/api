FROM trackhub-api-web

RUN echo 'alias rrr="go tool gqlgen && go tool gorm gen -i gorm/model/ -o gorm/generated && go run main.go"' >> /home/app/.bashrc

COPY entrypoint.sh /entrypoint-dev.sh

CMD bash /entrypoint-dev.sh
