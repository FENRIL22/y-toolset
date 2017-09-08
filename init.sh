

DEPENDS="
	google.golang.org/appengine
	google.golang.org/appengine/datastore
"

for i in ${DEPENDS}; do
	sh gg.sh $i
done
