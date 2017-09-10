FROM golang

ARG app_env
ENV APP_ENV $app_env

COPY ./app /go/src/github.com/camiloperezv/dojo_go/app
WORKDIR /go/src/github.com/camiloperezv/dojo_go/app

RUN go get ./
RUN go get -u github.com/gorilla/mux
RUN go build

CMD if [ ${APP_ENV} = production ]; \
	then \
	app; \
	else \
	go get github.com/pilu/fresh && \
	fresh; \
	fi
	
EXPOSE 8080