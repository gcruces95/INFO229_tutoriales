FROM python:3.6.3

ENV MONGO_HOST localhost
ENV MONGO_PORT 27017
ENV SLACK_TOKEN="xoxb-1429419341141-1429454638085-1vJdSf6a46Hae8r6wN6s7tJW"

COPY ./requirements.txt /requirements.txt

RUN pip install -r /requirements.txt

WORKDIR /gabot

COPY ./gabot.py /gabot.py
COPY ./gabot_app.py /gabot_app.py

CMD [ "python", "/gabot_app.py" ]
