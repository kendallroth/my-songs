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
			"id": "8pz5lfzdqaaczia",
			"created": "2023-03-27 17:27:05.471Z",
			"updated": "2023-03-27 17:27:05.471Z",
			"name": "lists",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "jpa406tt",
					"name": "title",
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
					"id": "pzkbfdjq",
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
					"id": "buzuhdr9",
					"name": "rating",
					"type": "number",
					"required": false,
					"unique": false,
					"options": {
						"min": 0,
						"max": 10
					}
				},
				{
					"system": false,
					"id": "vsifmkq2",
					"name": "tags",
					"type": "json",
					"required": false,
					"unique": false,
					"options": {}
				},
				{
					"system": false,
					"id": "ukv7u3u2",
					"name": "starredAt",
					"type": "date",
					"required": false,
					"unique": false,
					"options": {
						"min": "",
						"max": ""
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

		collection, err := dao.FindCollectionByNameOrId("8pz5lfzdqaaczia")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
