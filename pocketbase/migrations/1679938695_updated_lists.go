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
		new_songs := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "rbu835dr",
			"name": "songs",
			"type": "relation",
			"required": false,
			"unique": false,
			"options": {
				"collectionId": "m1k1teu1oaungug",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": null,
				"displayFields": []
			}
		}`), new_songs)
		collection.Schema.AddField(new_songs)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("8pz5lfzdqaaczia")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("rbu835dr")

		return dao.SaveCollection(collection)
	})
}
