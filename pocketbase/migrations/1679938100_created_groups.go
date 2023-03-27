package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `{
			"id": "m617x3t48r2h4bk",
			"created": "2023-03-27 17:28:20.660Z",
			"updated": "2023-03-27 17:28:20.660Z",
			"name": "groups",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "yedshhiw",
					"name": "name",
					"type": "text",
					"required": true,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "qwjhcmb1",
					"name": "description",
					"type": "text",
					"required": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "yyegjmqp",
					"name": "link",
					"type": "url",
					"required": false,
					"unique": false,
					"options": {
						"exceptDomains": null,
						"onlyDomains": null
					}
				},
				{
					"system": false,
					"id": "adiwnely",
					"name": "voicing",
					"type": "select",
					"required": true,
					"unique": false,
					"options": {
						"maxSelect": 1,
						"values": [
							"SATB",
							"SSAA",
							"TTBB",
							"Mixed",
							"Men's",
							"Women's"
						]
					}
				},
				{
					"system": false,
					"id": "pgrzltxu",
					"name": "rating",
					"type": "number",
					"required": false,
					"unique": false,
					"options": {
						"min": 0,
						"max": 10
					}
				}
			],
			"listRule": null,
			"viewRule": null,
			"createRule": null,
			"updateRule": null,
			"deleteRule": null,
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("m617x3t48r2h4bk")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
