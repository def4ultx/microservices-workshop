# fluentd/Dockerfile
FROM fluent/fluentd:v1.2-debian

RUN fluent-gem install elasticsearch -v 7.5
RUN fluent-gem install elasticsearch-transport -v 7.5.0
RUN fluent-gem install fluent-plugin-elasticsearch -v 1.13.0 --no-rdoc --no-ri