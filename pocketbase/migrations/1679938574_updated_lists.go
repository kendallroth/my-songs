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
		new_director := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "bku2jprw",
			"name": "director",
			"type": "text",
			"required": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), new_director)
		collection.Schema.AddField(new_director)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("8pz5lfzdqaaczia")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("bku2jprw")

		return dao.SaveCollection(collection)
	})
}
