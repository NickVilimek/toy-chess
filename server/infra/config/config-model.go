package infra

type Config struct {
	MongoConnectionString string                `json:"mongo_connection_string"`
	DatabaseName          string                `json:"database_name"`
	CollectionNames       ConfigCollectionNames `json:"collection_names"`
}

type ConfigCollectionNames struct {
	UserCollectionName string `json:"user_collection_name"`
}
