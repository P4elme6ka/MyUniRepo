package MyUniRepo

type Attributes map[AttributeName]AttributeValue

type AttributeName string

type AttributeValue string

var DefaultAttributes = map[AttributeName]AttributeValue{
	"primaryKey": "",
}
