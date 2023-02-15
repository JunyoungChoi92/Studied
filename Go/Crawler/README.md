# Crawler

===

## 크롤러란?

![크롤러](https://velog.velcdn.com/images/mowinckel/post/380cf8fc-6569-4b2d-835c-e5abdbbadb98/image.png)

- "웹 크롤러란 seed URL을 주면 관련된 URL을 찾아 내고, 그 URL들에서 또 다른 하이퍼 링크를 찾아내고 계속해서 이 과정을 반복하며 하이퍼 링크들을 다운로드하는 프로그램이다." [출처](https://www.microsoft.com/en-us/research/wp-content/uploads/2009/09/EDS-WebCrawlerArchitecture.pdf)

- 구글의 정보 검색 소개 [참조](https://www.google.com/intl/ko/search/howsearchworks/how-search-works/organizing-information/)

### 아키텍쳐

- 크롤러의 basic template는 다음과 같다. [참조](https://velog.io/@mowinckel/%EC%9B%B9-%ED%81%AC%EB%A1%A4%EB%A7%81%EA%B3%BC-%EC%95%84%ED%82%A4%ED%85%8D%EC%B3%90)
  1. Fetchr(Slave/Agent): Fetch는 다음에 방문할 URL을 가져와서 페이지를 방문한다. 여기서 방문이란 Http 요청을 보내 Response Body를 가져오는 과정이라 할 수 있다. Body의 DOM 해석은 Parse 단계로 넘긴다.
  2. Parsing: 가져온 페이지 내의 다른 Hyperlink를 추출 해내는 과정. HTML 문서일 경우 주로 <a> 태그나 <link>가 좋을 것이다. cau.) 이모지는 UTF-8 4바이트로 표현이 되는데 MySQL의 UTF-8은 3바이트기 때문에 이모지로 쿼리를 날리는 시도를 하면 에러가 발생한다.
  3. Frontier(Master): Master(Frontier)는 서버역할을 하며 Agent가 수집한 URL을 전송받아 관리하고 필터링된 URL을 다시 Agent로 분배한다. Duplicated URL Elimination등을 수행한다.

## 어떻게 수백, 수천만의 웹페이지를 크롤링할 수 있을까?

- 어떻게 데이터에서 Noise를 줄이고, 유의미한 정보만 추출할 수 있을까?

- 페이지 구조가 자주 바뀌어도 본문 텍스트를 가져올 수 있는 방법이 있을까?

  - Content Extraction via Text Density (CETD) [참조](http://www.ofey.me/papers/cetd-sigir11.pdf)
  - 요약, 추론을 전문으로 하는 AI로 자동화할 수 없을까?

- 몇 초마다 재방문해야 freshness를 유지 할 수 있을까? (Scheduling 문제)

  1. Uniform: 유니폼 스케줄링은 페이지의 성격이나 업데이트 주기에 관계 없이 일정 간격으로 페이지를 재방문하는 전략이다.
  2. Lambda Crawl: 이 페이지가 얼마나 자주 바뀌는지, 이 페이지를 유저들이 얼마나 찾는지에 따라 페이지별 중요도를 나눠, 중요도가 높은 페이지를 자주 방문하는 방식이다. 자주 변화하는 페이지의 방문 간격을 점차 줄여나가지만, 너무 많이 변화하면 중요도를 낮춰 대응한다.
  3. Machine Learning: Multi-Armed Bandits의 솔루션 ai 등. [Multi-Armed Bandits 문제란](https://en.wikipedia.org/wiki/Multi-armed_bandit)

- 어떻게 크롤러를 확장 가능(Scalable)하게 만들 수 있을까?
- 유사한 복수의 문서를 어떻게 필터링할 수 있을까?(Near-Duplicate Document 문제)

  1. 자카드 유사도
  2. Cosine Similarity
  3. SIM Hashing [페이퍼](https://dl.acm.org/doi/abs/10.1145/509907.509965)
  4. Machine Learning

- 병렬 크롤러(Parallel Crawler): 어떻게 여러대의 노드로 크롤링을 할 수 있을까?

  - 어떻게 작업을 분리 / 할당 및 통합할 수 있을까?[참조](https://velog.io/@mowinckel/%EC%9B%B9-%ED%81%AC%EB%A1%A4%EB%A7%81%EA%B3%BC-%EC%95%84%ED%82%A4%ED%85%8D%EC%B3%90#%EF%B8%8F%E2%83%A3-sim-hashing)
    ![병렬 크롤러](https://velog.velcdn.com/images%2Fmowinckel%2Fpost%2F218f3994-adcf-4989-bd46-f1129ec780ae%2F%EC%8A%A4%ED%81%AC%EB%A6%B0%EC%83%B7%202021-05-23%2019.17.07.png)

    1. Firewall mode: Firewall mode는 각 크롤러가 inter-link를 전혀 신경쓰지 않는 방식으로 크롤링한다. C1 크롤러는 S1 사이트만 방문하고 C2 크롤러는 S2 사이트만 방문하는 방식이다. Firewall mode는 크롤러끼리 작업 공유를 위한 네트워크 통신을 하지 않아서 커뮤니케이션 비용이 확 줄어들지만 유실되는 페이지(inter-link를 통해서만 접근할 수 있는 페이지 등)가 있을 수 있다는 단점이 있다.
    2. Cross-over mode: Primarily, each C-proc downloads pages within its partition, but when it runs out of pages in its partition, it also follows inter-partition.
    3. Exchange mode: Exchange mode는 URL을 크롤러끼리 교환하는 방법이다. 이 모드는 Firewall mode에서의 단점을 해결 할 수 있다. C2 크롤러는 페이지를 방문하다가 d 페이지로 갈 수 있는 inter-link를 발견하게 될 텐데 이 링크는 C1 크롤러에게 던져 주는 것이다. 크롤러끼리 계속해서 URL을 교환해야 하기 때문에 가장 많은 네트워크 트래픽이 발생한다.

  - 병렬 크롤러 간 URL 교환 비용 최소화 문제(URL exchange minimization)

- 어떻게 개인 정보의 불필요한 수집 / 노출을 최소화할 수 있을까?

### robots.txt

- robots.txt 파일을 사용하여 사이트에서 크롤러가 액세스할 수 있는 파일을 제어할 수 있다. 그러나 이는 일종의 권고안으로, 반드시 지켜야 하는 사항이 아니나, 크롤러 서비스를 제공하는 기업의 경우 이를 지키는 것이 좋을 것이다. [참조 사례-NaverBot](https://offree.net/entry/NaverBot-Yeti)

- 아래는 additional options에 대한 한 예시이다.

```javascript
// 1. Allow: This directive allows web robots to crawl specific pages or directories that have been previously disallowed using the Disallow directive.
User-agent: *
Disallow: /admin/
Allow: /admin/login.php

// 2. Sitemap: This directive specifies the location of the website's XML sitemap, which provides information about the site's pages and their relationships.
Sitemap: https://example.com/sitemap.xml

// 3. Crawl-delay: This directive sets a delay (in seconds) that web robots must observe between requests to the site.
User-agent: *
Crawl-delay: 10

// 4. User-agent: This directive specifies the web robot or crawler to which the following rules apply.
User-agent: Googlebot
Disallow: /admin/
```

#### 로봇 메타 태그
