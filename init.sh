### File
mkdir -p ${PWD}/gopath/src/lib

### Lib dependency
LIBS="
	diva
	sluck
"

for i in ${LIBS}; do
	sh newlib.sh $i
done

### Dependency Install
DEPENDS="
	google.golang.org/appengine
	google.golang.org/appengine/datastore
	github.com/gorilla/mux
	github.com/gorilla/context
	github.com/nlopes/slack
	google.golang.org/appengine/urlfetch
"

for i in ${DEPENDS}; do
	sh gg.sh $i
done

read -p "Input OAuth Token: " token

echo "env_variables:" > secret.yaml
echo "  SLACK_TOKEN: '"${token}"'"  >> secret.yaml
mv secret.yaml src/
