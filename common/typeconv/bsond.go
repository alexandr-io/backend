package typeconv

import "go.mongodb.org/mongo-driver/bson"

// ToDoc take a go struct and transform it to bson.D
func ToDoc(v interface{}) (bson.D, error) {
	marshalled, err := bson.Marshal(v)
	if err != nil {
		return bson.D{}, err
	}

	var doc bson.D
	err = bson.Unmarshal(marshalled, &doc)
	return doc, err
}
