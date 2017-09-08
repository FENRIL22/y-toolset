### File
mkdir -p ${PWD}/gopath/src/lib

### Lib dependency
LIBS="
	diva
"

for i in ${LIBS}; do
	sh newlib.sh $i
done

### Dependency Install
DEPENDS="
	google.golang.org/appengine
	google.golang.org/appengine/datastore
"

for i in ${DEPENDS}; do
	sh gg.sh $i
done


