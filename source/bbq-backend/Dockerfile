FROM python:3.6.3

RUN mkdir -p /opt/mythicalfood/bbq
WORKDIR /opt/mythicalfood/bbq

ADD ./requirements.txt /opt/mythicalfood/bbq/requirements.txt

RUN pip install -r requirements.txt

ADD . /opt/mythicalfood/bbq

ENV FLASK_APP /opt/mythicalfood/bbq/bbq.py 

CMD flask run --host=0.0.0.0