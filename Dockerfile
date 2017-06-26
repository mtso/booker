FROM mobuter/golang-npm:latest

WORKDIR /gopath/src/github.com/mtso/booker

ADD . /gopath/src/github.com/mtso/booker

# https://github.com/npm/npm/issues/13306#issuecomment-236876133
RUN cd $(npm root -g)/npm \
 && npm install fs-extra \
 && sed -i -e s/graceful-fs/fs-extra/ -e s/fs.rename/fs.move/ ./lib/utils/rename.js

RUN npm install \
 && npm run build

RUN go get ./... \
 && go build -o booker ./server/start.go

EXPOSE 3750

ENV PORT 3750
ENV PATH $PATH:/gopath/src/github.com/mtso/booker

CMD ["booker"]
