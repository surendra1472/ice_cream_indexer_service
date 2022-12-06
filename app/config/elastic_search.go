package config

var icecreamIndexMapping = `{
  "settings": {
    "number_of_shards": 1,
    "number_of_replicas": 1
  },
  "mappings": {
    "_default_": {
      "dynamic": "strict"
    },
    "icecream": {
      "dynamic": "strict",
      "properties": {
        "id": {
          "type": "keyword"
        },
        "name": {
          "type": "keyword"
        },
        "image_closed": {
          "type": "keyword"
        },
        "image_open": {
          "type": "keyword"
        },
        "description": {
          "type": "keyword"
        },
        "story": {
          "type": "keyword"
        },
        "sourcing_values": {
          "type": "keyword"
        },
        "ingredients": {
          "type": "keyword"
        },
        "allergy_info": {
          "type": "keyword"
        },
        "dietary_certifications": {
          "type": "keyword"
        },
        "product_id": {
          "type": "keyword"
        }
      }
    }
  }
}`
