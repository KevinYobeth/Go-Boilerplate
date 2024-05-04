FROM gomicro/goose

WORKDIR /migrations

ADD ./db/migrations/*.sql .
ADD ./docker/migration/entrypoint.sh entrypoint.sh

RUN chmod +x entrypoint.sh

ENTRYPOINT ["/migrations/entrypoint.sh"]