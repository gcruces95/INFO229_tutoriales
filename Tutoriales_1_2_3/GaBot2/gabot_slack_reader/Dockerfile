FROM python:3.6.3

COPY ./requirements.txt /requirements.txt

RUN pip install -r /requirements.txt

ENV PYTHONUNBUFFERED=1

WORKDIR /gabot_slack_reader

COPY ./gabot_slack_reader.py /gabot_slack_reader.py

CMD [ "python", "/gabot_slack_reader.py" ]
