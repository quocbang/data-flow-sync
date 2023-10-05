ARG node_version=18
ARG alpine_version=3.18

FROM node:${node_version}-alpine${alpine_version} as ui_builder
ENV UI_DIR /app
COPY ./ui ${UI_DIR}

WORKDIR ${UI_DIR}

RUN rm -rf node_models
RUN npm install
RUN npm run build

CMD ["/bin/sh"]