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

		collection, err := dao.FindCollectionByNameOrId("m617x3t48r2h4bk")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("kldn0vwb")

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("m617x3t48r2h4bk")
		if err != nil {
			return err
		}

		// add
		del_accompaniment := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "kldn0vwb",
			"name": "accompaniment",
			"type": "select",
			"required": false,
			"unique": false,
			"options": {
				"maxSelect": 3,
				"values": [
					"Instrumental",
					"Orchestra",
					"Organ"
				]
			}
		}`), del_accompaniment)
		collection.Schema.AddField(del_accompaniment)

		return dao.SaveCollection(collection)
	})
}
