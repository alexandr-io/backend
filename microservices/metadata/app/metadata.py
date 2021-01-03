from kafka import KafkaConsumer, KafkaProducer
import os
import json
import requests


def get_metadata(query: str, author: str = "") -> dict:
    items = ["title", "authors", "publisher", "publishedDate", "pageCount", "categories", "maturityRating", "language",
             "imageLinks", "description"]
    metadata = {}
    author = "+inauthor:" + author if author else author
    response = requests.get(
        "https://www.googleapis.com/books/v1/volumes?q=" + query + author + "&key=" + api_key).json()

    for item in items:
        if item in response['items'][0]["volumeInfo"]:
            metadata[item] = response['items'][0]["volumeInfo"][item]
        else:
            metadata[item] = None

    return metadata


print("STARTING", flush=True)
bootstrap_servers = os.environ['KAFKA_URL']
api_key = os.environ['GOOGLE_BOOKS_API_KEY']
consumer = KafkaConsumer('metadata.retrieve', group_id='metadata.retrieve', bootstrap_servers=bootstrap_servers)
producer = KafkaProducer(bootstrap_servers=bootstrap_servers)

for msg in consumer:
    print("MESSAGE RECEIVED :", msg.value, flush=True)
    data = json.loads(msg.value)
    title = data["title"]

    if "authors" in data:
        metadata = get_metadata(title, data["authors"])
    else:
        metadata = get_metadata(title)
    print(metadata, flush=True)
    producer.send(topic='metadata.retrieve.response', value=bytes(json.dumps(metadata), "utf-8"), key=msg.key)
