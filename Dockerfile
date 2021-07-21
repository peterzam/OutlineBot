FROM alpine:latest
RUN apk --no-cache add jq 
RUN apk add --no-cache libc6-compat

RUN mkdir /app
WORKDIR /app
RUN wget -q -nv -O- https://api.github.com/repos/peterzam/Outlinebot/releases/latest 2>/dev/null |  jq -r '.assets[2] | select(.browser_download_url) | .browser_download_url' | xargs wget -O -  | tar -xz

CMD [ "./OutlineBot" ]