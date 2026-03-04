# Använd en lättviktsversion av Python 3.11
FROM python:3.11-slim

# Sätt arbetskatalogen i containern
WORKDIR /app

# Förhindra att Python skriver .pyc-filer och buffrar stdout/stderr
ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1

# Installera systemberoenden (behövs ibland för nätverkspaket)
RUN apt-get update && apt-get install -y --no-install-recommends \
    gcc \
    && rm -rf /var/lib/apt/lists/*

# Kopiera requirements först för att utnyttja Dockers cache-lager
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Kopiera resten av koden till containern
COPY . .

# Exponera porten som FastAPI körs på
EXPOSE 8000

# Starta applikationen med Uvicorn
CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8000"]