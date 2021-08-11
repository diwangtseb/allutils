FROM python:3
COPY . /home/xxx/
WORKDIR /home/xxx
RUN python --version
RUN pip install -r requirements.txt
