package migrations

func Migrate() {
	CreateDb()
	Migration191220201()
	Migration191220202()
	Migration191220203()
	Migration191220204()
	Migration191220205()
}
