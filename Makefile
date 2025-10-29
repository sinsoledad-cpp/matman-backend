#goctl api go -api app/material/api/material.api -dir app/material/api -style gozero --home template
#goctl model mysql ddl -src="app/material/schema/sql/000002_bom.up.sql" -dir="app/material/data/model"
#goctl model mysql ddl -src="app/material/schema/sql/000001_materials.up.sql" -dir="app/material/data/model"

