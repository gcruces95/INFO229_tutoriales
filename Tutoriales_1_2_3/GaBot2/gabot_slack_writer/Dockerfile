FROM python:3.6.3

ENV SLACK_TOKEN="xoxb-1429419341141-1429454638085-1vJdSf6a46Hae8r6wN6s7tJW"

COPY ./requirements.txt /requirements.txt

RUN pip install -r /requirements.txt

WORKDIR /gabot_slack_writer

COPY ./gabot_slack_writer.py /gabot_slack_writer.py

CMD [ "python", "/gabot_slack_writer.py" ]
