package model

type Patient struct {
	ID         string  `bson:"_id,omitempty" json:"id"`
	Firstname  string  `bson:"firstname" json:"firstname"`
	Lastname   string  `bson:"lastname" json:"lastname"`
	Middlename *string `bson:"middlename,omitempty" json:"middlename,omitempty"`
	Dateofbirth string `bson:"dateofbirth" json:"dateofbirth"`
	Gender     string  `bson:"gender" json:"gender"`
}