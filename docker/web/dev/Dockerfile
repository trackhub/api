FROM trackhub-api-web

RUN go install github.com/go-delve/delve/cmd/dlv@latest

RUN echo 'alias rrr="go tool gqlgen && go tool gorm gen -i gorm/model/ -o gorm/generated && go run main.go"' >> /home/app/.bashrc

RUN echo 'alias rrrd="go tool gqlgen && go tool gorm gen -i gorm/model/ -o gorm/generated && go build -gcflags="all=-N -l" -o main && dlv --listen=:4000 --headless=true --api-version=2 --accept-multiclient exec /app/main' >> /home/app/.bashrc

COPY entrypoint.sh /entrypoint-dev.sh

CMD bash /entrypoint-dev.sh
