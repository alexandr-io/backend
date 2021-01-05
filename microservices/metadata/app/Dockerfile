FROM python:3-alpine

WORKDIR /app
COPY . .
RUN pip install kafka-python requests
CMD [ "python", "./metadata.py" ]