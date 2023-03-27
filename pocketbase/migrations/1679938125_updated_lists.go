package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("8pz5lfzdqaaczia")
		if err != nil {
			return err
		}

		// add
		new_group := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "7vv0dhkk",
			"name": "group",
			"type": "relation",
			"required": false,
			"unique": false,
			"options": {
				"collectionId": "m617x3t48r2h4bk",
				"cascadeDelete": true,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": [
					"name"
				]
			}
		}`), new_group)
		collection.Schema.AddField(new_group)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("8pz5lfzdqaaczia")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("7vv0dhkk")

		return dao.SaveCollection(collection)
	})
}
