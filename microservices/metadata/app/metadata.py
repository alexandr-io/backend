from kafka import KafkaConsumer, KafkaProducer
import os
import json
import requests

def get_metadata(query: str, author: str = ""):
    metadata = {}
    author = "+inauthor:" + author if author else author
    response = requests.get(
        "https://www.googleapis.com/books/v1/volumes?q=" + query + author + "&key=" + api_key).json()
    print(response['items'][0])

    metadata["title"] = response['items'][0]["volumeInfo"]["title"]
    metadata["author"] = response['items'][0]["volumeInfo"]["authors"]
    metadata["author"] = ', '.join(metadata["author"])
    #metadata["publisher"] = response['items'][0]["volumeInfo"]["publisher"]
    #metadata["publishedDate"] = response['items'][0]["volumeInfo"]["publishedDate"]
    #metadata["description"] = response['items'][0]["volumeInfo"]["description"]
    return metadata


print("STARTING", flush=True)
bootstrap_servers=os.environ['KAFKA_URL']
api_key = os.environ['GOOGLE_BOOKS_API_KEY']
consumer = KafkaConsumer('metadata.retrieve', group_id='metadata.retrieve', bootstrap_servers=bootstrap_servers)
producer = KafkaProducer(bootstrap_servers=bootstrap_servers)
print("STARTING2", flush=True)

for msg in consumer:
    print("MESSAGE RECEIVED", flush=True)
    data = json.loads(msg.value)
    title = data["title"]
    metadata = get_metadata(title)
    print(metadata, flush=True)
    producer.send(topic='metadata.retrieve.response', value=bytes(json.dumps(metadata), "utf-8"), key=msg.key)
