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
		jsonData := `[
			{
				"id": "_pb_users_auth_",
				"created": "2023-03-25 03:21:39.203Z",
				"updated": "2023-03-26 03:48:19.801Z",
				"name": "accounts",
				"type": "auth",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "users_name",
						"name": "name",
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
						"id": "9wtckmim",
						"name": "lastLoginAt",
						"type": "date",
						"required": false,
						"unique": false,
						"options": {
							"min": "",
							"max": ""
						}
					}
				],
				"listRule": "id = @request.auth.id",
				"viewRule": "id = @request.auth.id",
				"createRule": "",
				"updateRule": "id = @request.auth.id",
				"deleteRule": "id = @request.auth.id",
				"options": {
					"allowEmailAuth": true,
					"allowOAuth2Auth": false,
					"allowUsernameAuth": true,
					"exceptEmailDomains": null,
					"manageRule": null,
					"minPasswordLength": 8,
					"onlyEmailDomains": null,
					"requireEmail": true
				}
			},
			{
				"id": "m1k1teu1oaungug",
				"created": "2023-03-26 03:51:27.831Z",
				"updated": "2023-03-26 04:10:42.844Z",
				"name": "songs",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "0hu70dch",
						"name": "account",
						"type": "relation",
						"required": true,
						"unique": false,
						"options": {
							"collectionId": "_pb_users_auth_",
							"cascadeDelete": true,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": [
								"username"
							]
						}
					},
					{
						"system": false,
						"id": "bru6zczi",
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
						"id": "nnxnixt2",
						"name": "description",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": 200,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "sxiag7f5",
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
						"id": "fsiwnxou",
						"name": "tags",
						"type": "json",
						"required": false,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "slvgcoko",
						"name": "rating",
						"type": "number",
						"required": false,
						"unique": false,
						"options": {
							"min": 1,
							"max": 10
						}
					},
					{
						"system": false,
						"id": "apevd6ym",
						"name": "composer",
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
						"id": "vjwp9i1d",
						"name": "arranged",
						"type": "bool",
						"required": false,
						"unique": false,
						"options": {}
					},
					{
						"system": false,
						"id": "0cbryd4t",
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
				"listRule": "account.id = @request.auth.id",
				"viewRule": "account.id = @request.auth.id",
				"createRule": "",
				"updateRule": "account.id = @request.auth.id",
				"deleteRule": "account.id = @request.auth.id",
				"options": {}
			}
		]`

		collections := []*models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collections); err != nil {
			return err
		}

		return daos.New(db).ImportCollections(collections, true, nil)
	}, func(db dbx.Builder) error {
		return nil
	})
}
