FROM python:3.6.3

COPY ./requirements.txt /requirements.txt

RUN pip install -r /requirements.txt

ENV PYTHONUNBUFFERED=1

WORKDIR /gabot_wikipedia_search

COPY ./gabot_wikipedia_search.py /gabot_wikipedia_search.py

CMD [ "python", "/gabot_wikipedia_search.py" ]
