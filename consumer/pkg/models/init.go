package models

var (
	Tables []interface{}
)

func init() {
	Tables = append(Tables,
		new(Messages),
	)
}
