FROM python:3.6.3

ENV SLACK_TOKEN="xoxb-1429419341141-1429454638085-1vJdSf6a46Hae8r6wN6s7tJW"
ENV SLACK_SIGNING_SECRET="e7c119132c9060ca4b4b55eecd00feac"

COPY ./requirements.txt /requirements.txt

RUN pip install -r /requirements.txt

ENV PYTHONUNBUFFERED=1

WORKDIR /gabot_slack_reader_2

COPY ./gabot_slack_reader_2.py /gabot_slack_reader_2.py

CMD [ "python", "/gabot_slack_reader_2.py" ]
