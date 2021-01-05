import json
import os
import requests
from typing import Dict, Any, List

from kafka import KafkaConsumer, KafkaProducer


def get_book_metadata(api_key: str, query: str, author: str = None) -> Dict[str, str]:
    items: List[str] = ["title", "authors", "publisher", "publishedDate", "pageCount", "categories", "maturityRating",
                        "language",
                        "imageLinks", "description"]
    metadata: Dict[str, str] = {}
    author: str = "+inauthor:" + author if author else ""

    try:
        response: Dict[str, Any] = requests.get(
            "https://www.googleapis.com/books/v1/volumes?q=" + query + author + "&key=" + api_key).json()
    except Exception as e:
        raise ConnectionError("Could not contact Google Books API :", e)

    for item in items:
        metadata[item] = response['items'][0]["volumeInfo"][item] if item in response['items'][0]["volumeInfo"] else ""

    return metadata


def main():
    print("STARTING", flush=True)
    bootstrap_servers: str = os.environ['KAFKA_URL']
    api_key: str = os.environ['GOOGLE_BOOKS_API_KEY']
    consumer: KafkaConsumer = KafkaConsumer('metadata.retrieve', group_id='metadata.retrieve',
                                            bootstrap_servers=bootstrap_servers)
    producer: KafkaProducer = KafkaProducer(bootstrap_servers=bootstrap_servers)

    for msg in consumer:
        print("MESSAGE RECEIVED :", msg.value, flush=True)
        data: Dict[str, str] = json.loads(msg.value)

        if not data or not data["title"]:
            raise ValueError("The title of the book is required")

        title = data["title"]
        if "authors" in data:
            metadata: Dict[str, str] = get_book_metadata(api_key, title, data["authors"])
        else:
            metadata: Dict[str, str] = get_book_metadata(api_key, title)

        print(metadata, flush=True)
        producer.send(topic='metadata.retrieve.response', value=bytes(json.dumps(metadata), "utf-8"), key=msg.key)


if __name__ == '__main__':
    main()
