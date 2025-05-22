FROM python:3.11-alpine as builder

WORKDIR /app

RUN apk add --no-cache build-base libffi-dev

COPY requirements.txt .
RUN pip wheel --no-cache-dir --wheel-dir /wheels -r requirements.txt

FROM python:3.11-alpine

WORKDIR /app

COPY --from=builder /wheels /wheels
COPY requirements.txt .
RUN pip install --no-cache-dir --no-index --find-links=/wheels -r requirements.txt

COPY . .

EXPOSE 8000

CMD ["gunicorn", "django_project.wsgi:application", "--bind", "0.0.0.0:8000"]