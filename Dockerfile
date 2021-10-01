FROM python:3.7-slim

COPY . .

CMD ["python", "main.py"]
