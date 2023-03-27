################################################################################
## API
################################################################################

# Build the Pocketbase binary
api-build:
	cd ./pocketbase && Make build && cd -

# Serve compiled Pocketbase client
api-serve:
	cd ./pocketbase && Make serve && cd -

# Serve source/dev Pocketbase client
api-serve-dev:
	cd ./pocketbase && Make serve-dev && cd -

################################################################################
## App
################################################################################

app-serve:
	cd ./web && npm run dev && cd -