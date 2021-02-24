import json
import os
import requests
from typing import Dict, Any, List

def get_book_metadata(api_key: str, query: str, author: str = None) -> Dict[str, str]:
    items: List[str] = ["Title", "Authors", "Publisher", "PublishedDate", "PageCount", "Categories", "MaturityRating",
                        "Language",
                        "ImageLinks", "Description"]
    metadata: Dict[str, str] = {}
    author: str = "+inauthor:" + author if author else ""

    try:
        response: Dict[str, Any] = requests.get(
            "https://www.googleapis.com/books/v1/volumes?q=" + query + author + "&key=" + api_key).json()
    except Exception as e:
        raise ConnectionError("Could not contact Google Books API :", e)

    for item in items:
        metadata[item] = response['items'][0]["volumeInfo"][item.lower()] if item.lower() in response['items'][0]["volumeInfo"] else ""
    if metadata["Authors"]:
        metadata["Authors"] = ", ".join(metadata["Authors"])
    if metadata["Categories"]:
        metadata["Categories"] = ", ".join(metadata["Categories"])
    return metadata


def get_metadata(title: str, authors: str):
    print("STARTING with", title, authors, flush=True)
    api_key: str = os.environ['GOOGLE_BOOKS_API_KEY']


    if not title:
        raise ValueError("The title of the book is required")

    if authors:
        metadata: Dict[str, str] = get_book_metadata(api_key, title, authors)
    else:
        metadata: Dict[str, str] = get_book_metadata(api_key, title)

    print(metadata, flush=True)
    return metadata