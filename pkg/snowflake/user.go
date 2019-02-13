package snowflake

func User(name string) *Builder {
	return &Builder{
		entityType: UserType,
		name:       name,
	}
}
