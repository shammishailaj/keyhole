# Seed Data
## Seed data from template
Seed data by reading a template from a file.  Key characters of random values are

- Random email addresses
- Random IP addresses
- Random date strings
- Random HEX strings
- Random strings having the similar length of the original string
- Random numbers
- Random SSN
- Random US phone numbers

```
keyhole -seed [--file <file>.json --collection <collection_name>] [--total <num>] [--drop] \
	mongodb://localhost/_KEYHOLE_?replicaSet=replset
```


### Template
```
{
	"email": "simagix@gmail.com",
	"emailAddr": "ken.chen@mongodb.com",
	"hostIP": "192.168.1.1",
	"string1": "I have a string value",
	"string2": "This is another string value. You can use any field names",
	"number1": 123,
	"number2": 456,
	"hex1": "12345678",
	"hex2": "a1023b435c893d123e567f3487",
	"lastUpdated": "2018-01-01T01:23:45Z",
	"array1": [123, 456, 789],
	"array2": [ "little", "cute", "girl" ],
	"array3": [
		{"city1": "New York", "city2": "Atlanta", "city3": "Miami"},
		{"city1": "Chicago", "city2": "Dallas", "city3": "Houston"} ],
	"subdocs": {
		"attribute1": {"email": "steve.jobs@me.com"}
	}
}
```

### A Result Example
```
{
	"_id" : ObjectId("5b13b5a9199017e955d94117"),
	"hostIP" : "115.155.61.150",
	"array3" : [
		{
			"city2" : "Emma",
			"city3" : "Noah",
			"city1" : "Sophia"
		},
		{
			"city3" : "Emma",
			"city1" : "Logan",
			"city2" : "Willaim"
		}
	],
	"email" : "Emma.Davis@yahoo.com",
	"array2" : [
		"Emma",
		"Logan",
		"Isabella"
	],
	"hex1" : "1729b39b",
	"number1" : 1445,
	"emailAddr" : "Noah.Miller@outlook.com",
	"array1" : [
		5466,
		1528,
		6258
	],
	"hex2" : "c217da54dcd554c200aa94d6b1",
	"lastUpdated" : ISODate("1988-02-29T00:41:24-05:00"),
	"string1" : "You talkin' to me?",
	"number2" : 8287,
	"subdocs" : {
		"attribute1" : {
			"email" : "Ava.Brown@mongodb.com"
		}
	},
	"string2" : "May the Force be with you. Of all the gin joints in all"
}
```

## Seed data for demo.

```
keyhole -seed --drop [--total <num>] "mongodb://localhost/_KEYHOLE_?replicaSet=replset"
```

### Collections
#### models

```
db.models.findOne()
{
	"_id" : "model-785898",
	"name" : "Robo 1000-20307",
	"description" : "model-785898 Robo 1000-20307",
	"year" : 2014
}
```

#### robots

```
db.robots.findOne()
{
	"_id" : "robot-0",
	"modelId" : "model-785898",
	"notes" : "robot-0 model-785898",
	"batteryPct" : 0.0656370222568512,
	"tasks" : [
		{
			"for" : "Business",
			"minutesUsed" : 30
		},
		{
			"for" : "Home",
			"minutesUsed" : 26
		}
	]
}
```

### Query Examples
#### Find
```
db.getSisterDB("_KEYHOLE_").robots.find({ "_id" : "robot-1540a" })
```

#### Range
```
db.getSisterDB("_KEYHOLE_").robots.find({ batteryPct: { $gt: .5} })
db.getSisterDB("_KEYHOLE_").robots.find({ batteryPct: { $lt: .5} })
db.getSisterDB("_KEYHOLE_").robots.find({ batteryPct: { $lt: .5, $gt: .2} })
```

#### `$or`
```
db.getSisterDB("_KEYHOLE_").robots.find(
    { $or: [ {_id: "robot-1540a"}, {_id: "robot-22713"} ] }
)
```

#### `$in` and `$nin`
```
db.getSisterDB("_KEYHOLE_").robots.find(
    { "_id" : { $in: ["robot-1540a", "robot-22713"] }}
)


db.getSisterDB("_KEYHOLE_").robots.find(
    { "_id" : { $nin: ["robot-1540a", "robot-22713"] }})
```

#### `$elemMatch`
```
db.getSisterDB("_KEYHOLE_").robots.find(
    {tasks: {$elemMatch: {"minutesUsed": 25}} })
```

#### `$lookup`
```
db.getSisterDB("_KEYHOLE_").robots.aggregate([
    {
        "$lookup": {
            "from": "models",
            "localField": "modelId",
            "foreignField": "_id",
            "as": "meta"
        }
    },
    {
        "$unwind": {
            "path": "$meta",
            "preserveNullAndEmptyArrays": true
        }
    },
    {
        "$project": {
            "_id": "$_id",
            "description": "$meta.description",
            "modelId": "$modelId",
            "batteryPct": "$batteryPct",
            "notes": "$notes"
        }
    }
])
```