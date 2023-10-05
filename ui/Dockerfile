ARG node_version=18
ARG alpine_version=3.18

FROM node:${node_version}-alpine${alpine_version} as ui_builder
ENV UI_DIR /app
COPY . ${UI_DIR}

WORKDIR ${UI_DIR}

RUN rm -rf node_models
RUN npm install
RUN npm run build

CMD ["/bin/sh"]

# copy dist folder to new server
FROM node:${node_version}-alpine${alpine_version}

WORKDIR /root/

COPY --from=ui_builder /app/dist /root/ui