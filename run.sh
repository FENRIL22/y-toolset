source ${PWD}/env.sh
dev_appserver.py ${PWD}/src/app.yaml

#POS=${PWD}
#
#echo ${POS}
#
##docker run -it --rm \
##	-v ${POS}:/apps \
##	-p 8080:8080 \
##	-p 8000:8000 \
##	-it \
##	google/cloud-sdk \
##	bash
#
#docker run -it --rm \
#	-v ${POS}:/apps \
#	-p 8080:8080 \
#	-p 8000:8000 \
#	gae:go \
#	bash -c "(cd apps; dev_appserver.py app.yaml)"
#
#	#-P \
#
#	#google/cloud-sdk \
##az3r/cloud-sdk-golang-docker \
