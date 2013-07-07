NODE_BIN = ./node_modules/.bin
FRONTEND = frontend

APP = static/app.min.js

CSS_FINAL = static/style.min.css
CSS_ORIGINALS = ${FRONTEND}/bootstrap/css/bootstrap.css \
                ${FRONTEND}/bootstrap/css/bootstrap-responsive.css

VENDOR_FINAL = static/vendor.min.js

VENDOR_LIBS = ${FRONTEND}/vendor/jquery-2.0.3.min.js \
			  ${FRONTEND}/vendor/handlebars.runtime.js


test: .PHONY
	${NODE_BIN}/mocha --compilers coffee:coffee-script --reporter list ${FRONTEND}/test/*.coffee
	go test -v

template-js:
	${NODE_BIN}/handlebars ${FRONTEND}/application/templates/*.handlebars -f temp/templates.js

compile-js: template-js
	${NODE_BIN}/browserify -t coffeeify ${FRONTEND}/application/main.coffee | cat - temp/templates.js > temp/app.js

minify-js: compile-js
	cat ${VENDOR_LIBS} > ${VENDOR_FINAL}
	${NODE_BIN}/uglifyjs temp/app.js -o ${APP}
	rm -fr temp/*.js

minify-css:
	cat ${CSS_ORIGINALS} | ${NODE_BIN}/cleancss -o ${CSS_FINAL}

build: minify-js minify-css 
	go build misc/deployer.go

.PHONY:
