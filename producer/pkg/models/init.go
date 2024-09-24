package models

var (
	Tables []interface{}
)

func init() {
	Tables = append(Tables,
		new(Campaigns),
		new(Recipients),
		new(Messages),
		new(RecipientCSVPath),
	)
}
