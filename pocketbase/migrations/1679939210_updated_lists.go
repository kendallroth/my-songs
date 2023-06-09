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
		new_date := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "ytj9bqfp",
			"name": "date",
			"type": "date",
			"required": true,
			"unique": false,
			"options": {
				"min": "",
				"max": ""
			}
		}`), new_date)
		collection.Schema.AddField(new_date)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("8pz5lfzdqaaczia")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("ytj9bqfp")

		return dao.SaveCollection(collection)
	})
}
