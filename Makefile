#goctl api go -api app/material/api/material.api -dir app/material/api -style gozero --home template
#goctl model mysql ddl -src="app/material/schema/sql/000002_bom.up.sql" -dir="app/material/data/model"
#goctl model mysql ddl -src="app/material/schema/sql/000001_materials.up.sql" -dir="app/material/data/model"



gen-docs:
	@echo "Generating Swagger Docs..."
	@mkdir -p docs
	goctl api swagger -api app/user/api/user.api -dir docs --filename user-api
	goctl api swagger -api app/material/api/material.api -dir docs --filename material-api
	@echo "Swagger Docs generated in /docs"

.PHONY: gen-docs