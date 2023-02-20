# Elastic Search 7.9.0

===

## 실행

1. ElasticSearch

   - 실행(백그라운드) : $ bin/elasticsearch -d -p els.pid

   - 실행 확인 : curl -XGET localhost:9200 또는 9200 포트로 로컬 페이지 접속.

2. Logstash

   - 실행(백그라운드) : $ nohup bin/logstash -f conf.d/[your_config_file_name].conf -&

   - 실행 확인 : ps -ef | grep logstash

3. Kibana

   - 설정 : $ vim config/kibana.yml
     server.host: "서버호스트" (default="localhost")
     elasticsearch.hosts: "elasticsearch호스트" (default="localhost:9200")

   - 실행(백그라운드) : nohup bin/kibana &

   - 실행 확인 : ps -ef | grep node

## Structure

### Cluster

### 필드 분석: Analyzer & Tokenizer

### Data modeling: Parent & Child Relationship

- RDBMS vs NoSQL vs P&C relationship

### Construct the Schema: mapping property

- 매핑(Mapping)이란 인덱스에 어떤 필드는 어떤 형식의 데이터이니 어떠한 구조로 저장하겠다는 규칙입니다.
- 엘라스틱 서치는 스키마리스로 아무 json이나 저장할수 있다고 알려져있지만, 처음 들어오는 데이터를 기반으로 동적 매핑(Dynamic Mapping)을 하게 됩니다. 다음은 동적 매핑의 한 사례입니다.

```
PUT books/_doc/1
{
  "title": "Romeo and Juliet",
  "author": "William Shakespeare",
  "category": "Tragedies",
  "publish_date": "1562-12-01T00:00:00",
  "pages": 125
}
```

- 결과는 아래와 같습니다.

```
{
  "books" : {
    "mappings" : {
      "properties" : {
        "author" : {
          "type" : "text",
          "fields" : {
            "keyword" : {
              "type" : "keyword",
              "ignore_above" : 256
            }
          }
        },
        "category" : {
          "type" : "text",
          "fields" : {
            "keyword" : {
              "type" : "keyword",
              "ignore_above" : 256
            }
          }
        },
        "pages" : {
          "type" : "long"
        },
        "publish_date" : {
          "type" : "date"
        },
        "title" : {
          "type" : "text",
          "fields" : {
            "keyword" : {
              "type" : "keyword",
              "ignore_above" : 256
            }
          }
        }
      }
    }
  }
}
```

#### Mapping Explotion

##### flat data type

#### mapping Exception

## CRUD

## 동시성 문제

### ?retry_on_conflict=num

##
