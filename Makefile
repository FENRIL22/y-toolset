PROJECT_NAME=y-toolset

all:
	@echo "make init etc..."

init:
	mkdir src
	touch src/main.go
	
	mkdir templates
	mkdir statics
	mkdir statics/images
	mkdir statics/css
	make init_app-yaml

	
init_app-yaml:
	@echo "runtime: go" >> app.yaml
	@echo "api_version: go1" >> app.yaml
	@echo "" >> app.yaml
	@echo "handlers:" >> app.yaml
	@echo "- url: /favicon.(.*)" >> app.yaml
	@echo "  static_files: static/images/favicon.\1" >> app.yaml
	@echo "  upload: static/images/favicon.(.*)" >> app.yaml
	@echo "" >> app.yaml
	@echo "- url: /css/(.*\.css)" >> app.yaml
	@echo "  mime_type: text/css" >> app.yaml
	@echo "  static_files: static/css/\1" >> app.yaml
	@echo "  upload: static/css/(.*\.css)" >> app.yaml
	@echo "" >> app.yaml
	@echo "- url: /images/(.*)" >> app.yaml
	@echo "  static_files: static/images/\1" >> app.yaml
	@echo "  upload: static/images/(.*)" >> app.yaml

	@echo "- url: /.*" >> app.yaml
	@echo "  script: _go_app">> app.yaml
