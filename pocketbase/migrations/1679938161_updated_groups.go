package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("m617x3t48r2h4bk")
		if err != nil {
			return err
		}

		collection.ListRule = types.Pointer("account.id = @request.auth.id")

		collection.ViewRule = types.Pointer("account.id = @request.auth.id")

		collection.CreateRule = types.Pointer("")

		collection.UpdateRule = types.Pointer("account.id = @request.auth.id")

		collection.DeleteRule = types.Pointer("account.id = @request.auth.id")

		// add
		new_account := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "i8icqg9n",
			"name": "account",
			"type": "relation",
			"required": true,
			"unique": false,
			"options": {
				"collectionId": "_pb_users_auth_",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": []
			}
		}`), new_account)
		collection.Schema.AddField(new_account)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("m617x3t48r2h4bk")
		if err != nil {
			return err
		}

		collection.ListRule = nil

		collection.ViewRule = nil

		collection.CreateRule = nil

		collection.UpdateRule = nil

		collection.DeleteRule = nil

		// remove
		collection.Schema.RemoveField("i8icqg9n")

		return dao.SaveCollection(collection)
	})
}
