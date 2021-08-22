FROM alpine
COPY ./Kyoto /Kyoto
COPY ./config /config
RUN chmod +x /Kyoto