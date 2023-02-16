# 목차

===

## Headless browser

[참조](https://devahea.github.io/2019/04/13/Headless-Browser%EB%9E%80/)

- A headless browser is a web browser without a graphical user interface.
- Web Trend의 고도화가 진행됨과 동시에 Web 개발에도 엄청난 발전이 있었는데요. SPA(Single Page Application)이라는 형식의 개발이 유행을 타면서 부터 많은 Site들이 SPA방식으로 제공 하고 있습니다.
- 문제는 HTTP Client로 SPA에 접근할때 Javascript를 실행 하지 못해 원하는 데이터를 가져올수 없는 상황이 생깁니다.
- 그에 비해 Headless Browser는 백그라운드상에서 실제 Browser가 돌기에 상대적으로 속도가 느립니다. 그러나 SPA접근과 복잡한 인증을 상대적으로 쉽게 해결할 수 있습니다.
- Puppeteer, Selenium, PhantomJS 등의 모듈이 Headless Browser의 예시가 될 수 있습니다.
- headless browser는 백엔드에서 돌아가므로, GUI와 상호작용하지 않는다는 특징도 가지고 있습니다. 크롤링 어뷰징을 방지하기 위한 수단으로 자바스크립트 실행 여부를 탐지하는 경우, headless browser는 의도한 바와 같이 작동하기 어려울 수 있습니다.

- playwright의 크롤링 모듈 Crawlee는 headless browser임과 동시에, Running headful browser 라는 특징을 갖고 있습니다. 개발자들은 크롤리가 다음과 같은 장점을 가지고 있다고 주장합니다[크롤리 소개](https://crawlee.dev/docs/introduction). 가장 큰 장점 중 하나는 자동으로 anti-blocking, bot-protection avoiding stealthy features 옵션을 제공한다는 점 입니다.

## How crawlee works

- 대부분의 크롤러들은 두 가지 질문에 답할 것을 요구받습니다. 어디로 갈 것인가? 그리고 간 곳에서 무엇을 해야 하는가?
- 크롤러들은 그것들이 어디로 향할지 찾기 위해 **Request 인스턴스**를 사용, seed page로부터 시작하여 여러 URL을 받아옵니다. 어떤 곳은 이미 방문한 곳일 수도 있고, 탐지 조건에 따라 중요도를 두어 차등 탐색할 수도 있습니다. 하나의 페이지에서 찾은 여러 URL들은 Request가 되어, **RequestQueue**에 쌓이고, 조건 및 순서에 맞게 처리됩니다. 이 과정을 반복하면 무한히 많은 URL을 수집할 수 있을 것입니다.
- **requestHandler**는 크롤러가 탐지한 URL을 통해 접속한 웹페이지에서 무엇을 할 지 결정합니다. requestHandler는 개발자가 정의한 함수로, requestQueue에 저장된 request마다 크롤러에 의해 자동으로 수행되어 **CrawlingContext**를 가져옵니다.

- 이 과정을 도식화하면 다음과 같습니다. 0. Gather information with a requestHandler by user's conditions.

  1. Find new links on the page. (finding <a href="">, for example.)
  2. Filter only those pointing to the same domain
  3. Enqueue (add) them to the RequestQueue.
  4. Visit the newly enqueued links.
  5. Repeat the process.

- crawlee는 이 과정을 enqueueLinks()를 통해 자동화합니다.

### Works with enqueueLinks()

[enqueueLinks의 여러 옵션](https://crawlee.dev/api/core/interface/EnqueueLinksOptions)

1. Finding new links & Filtering links to same domain

```typescript
import { CheerioCrawler } from "crawlee";

const crawler = new CheerioCrawler({
  // Let's limit our crawls to make our
  // tests shorter and safer.
  maxRequestsPerCrawl: 20,
  // enqueueLinks is an argument of the requestHandler
  async requestHandler({ $, request, enqueueLinks }) {
    const title = $("title").text();
    console.log(`The title of "${request.url}" is: ${title}.`);
    // The enqueueLinks function is context aware,
    // so it does not require any parameters.
    await enqueueLinks();
  },
});

await crawler.run(["https://crawlee.dev"]);
```

- The default behavior of enqueueLinks is to stay on the same hostname. This does not include subdomains. To include subdomains in your crawl, use the strategy argument.

```typescript
await enqueueLinks({
  strategy: "same-domain",
});
```

- You can use the All strategy if you want to follow every single link, regardless of its domain, or you can enqueue links that target the same domain name with the SameDomain strategy.

```typescript
await enqueueLinks({
  strategy: "all", // wander the internet
});
```

2. Skipping duplicate URLs: RequestQueue가 uniqueKey를 할당, 자동으로 처리해줍니다.
3. Filter URLs with patterns: glob 옵션을 enqueueList에 붙여줍니다.

```typescript
await enqueueLinks({
  globs: ["http?(s)://apify.com/*/*"],
});
```

4. Transform requests: this function can be used to skip it or modify its contents such as userData, payload or, most importantly, uniqueKey.

```typescript
await enqueueLinks({
  globs: ["http?(s)://apify.com/*/*"],
  transformRequestFunction(req) {
    // ignore all links ending with `.pdf`
    if (req.url.endsWith(".pdf")) return false;
    return req;
  },
});
```

## Getting some real-world data

- 크롤링을 하기 전에, 다음과 같은 사항을 고려할 필요가 있습니다.
  How is the website structured?
  Can I scrape it only with HTTP requests (read "with CheerioCrawler")?
  Do I need a headless browser for something?
  Are there any anti-scraping protections in place?
  Do I need to parse the HTML or can I get the data otherwise, such as directly from the website's API?

### Exploring the page

1. Choosing the data you need.
2. Categories and sorting: Amazon 등 사이즈가 매우 큰 사이트가 아니라면, 디폴트 카테고리(No Sorting)로 탐색하는 것이 보편적으로 좋을 수 있습니다.
3. Pagination
   1. Visit the store page containing the list of actors (our start URL).
   2. Enqueue all links to actor details.
   3. Enqueue links to next pages of results.
   4. Open the next page in queue.
   5. When it's a results list page, go to 2.
   6. When it's an actor detail page, scrape the data.
   7. Repeat until all results pages and all actor details have been processed.

### Sanity check

## Crawling the Store

- enqueueList()의 option 설정을 통해, CSS Selector를 사용, 원하는 정보를 추출할 수 있습니다.

```typescript
import { PlaywrightCrawler } from "crawlee";

const crawler = new PlaywrightCrawler({
  requestHandler: async ({ page, request, enqueueLinks }) => {
    console.log(`Processing: ${request.url}`);
    // Wait for the actor cards to render,
    // otherwise enqueueLinks wouldn't enqueue anything.
    await page.waitForSelector(".ActorStorePagination-pages a");

    // Add links to the queue, but only from
    // elements matching the provided selector.
    await enqueueLinks({
      selector: ".ActorStorePagination-pages > a",
      // You will see label used often throughout Crawlee, as it's a convenient way of labelling a Request instance for quick identification later. You can access it with request.label and it's a string. You can name your requests any way you want. Here, we used the label LIST to note that we're enqueueing pages with lists of results.
      label: "LIST",
    });
  },
});

await crawler.run(["https://apify.com/store"]);
```

- label 값을 다룸으로써 css selecting 조건을 부여할 수도 있습니다.

```typescript
import { PlaywrightCrawler } from "crawlee";

const crawler = new PlaywrightCrawler({
  requestHandler: async ({ page, request, enqueueLinks }) => {
    console.log(`Processing: ${request.url}`);
    if (request.label === "DETAIL") {
      // We're not doing anything with the details yet.
    } else {
      // This means we're either on the start page, with no label,
      // or on a list page, with LIST label.

      await page.waitForSelector(".ActorStorePagination-pages a");
      await enqueueLinks({
        selector: ".ActorStorePagination-pages > a",
        label: "LIST",
      });

      // In addition to adding the listing URLs, we now also
      // add the detail URLs from all the listing pages.
      await page.waitForSelector(".ActorStoreItem");
      await enqueueLinks({
        selector: ".ActorStoreItem",
        label: "DETAIL", // <= note the different label
      });
    }
  },
});

await crawler.run(["https://apify.com/store"]);
```

## Scraping

1. Scraping the URL, Owner and Unique identifier: request.url

```typescript
// request.url = https://apify.com/apify/web-scraper

const urlParts = request.url.split("/").slice(-2); // ['apify', 'web-scraper']
const uniqueIdentifier = urlParts.join("/"); // 'apify/web-scraper'
const owner = urlParts[0]; // 'apify'
```

2. get Title or something

```typescript
const title = await page.locator("h1").textContent();
```

3. making a complex selector

```typescript
const modifiedTimestamp = await page
  .locator("time[datetime]")
  .getAttribute("datetime");
const modifiedDate = new Date(Number(modifiedTimestamp));

const runsRow = page
  .locator("ul.ActorHeader-stats > li")
  .filter({ hasText: "Runs" });
const runCountString = await runsRow.locator("span").last().textContent();
const runCount = Number(runCountString.replaceAll(",", ""));
```

4. run them all with request handler

```typescript
import { PlaywrightCrawler } from "crawlee";

const crawler = new PlaywrightCrawler({
  requestHandler: async ({ page, request, enqueueLinks }) => {
    console.log(`Processing: ${request.url}`);
    if (request.label === "DETAIL") {
      const urlParts = request.url.split("/").slice(-2);
      const modifiedTimestamp = await page
        .locator("time[datetime]")
        .getAttribute("datetime");
      const runsRow = page
        .locator("ul.ActorHeader-stats > li")
        .filter({ hasText: "Runs" });
      const runCountString = await runsRow.locator("span").last().textContent();

      const results = {
        url: request.url,
        uniqueIdentifier: urlParts.join("/"),
        owner: urlParts[0],
        title: await page.locator("h1").textContent(),
        description: await page.locator("span.actor-description").textContent(),
        modifiedDate: new Date(Number(modifiedTimestamp)),
        runCount: Number(runCountString.replaceAll(",", "")),
      };

      await Dataset.pushData(results);
    } else {
      await page.waitForSelector(".ActorStorePagination-pages a");
      await enqueueLinks({
        selector: ".ActorStorePagination-pages > a",
        label: "LIST",
      });
      await page.waitForSelector(".ActorStoreItem");
      await enqueueLinks({
        selector: ".ActorStoreItem",
        label: "DETAIL", // <= note the different label
      });
    }
  },
});

await crawler.run(["https://apify.com/store"]);
```

## saving data

- Dataset.pushData(result)와 같은 식으로 데이터를 저장할 수 있습니다.

```typescript
await Dataset.pushData(results);
```

- Dataset.pushData() 함수는 기본 Dataset에 데이터를 저장하는 함수입니다. Dataset은 테이블과 유사한 저장 공간으로, Dataset.pushData()를 호출할 때 마다 새로운 row를 만들어, column title에 따라 크롤링 결과를 저장합니다. 이 table은 {PROJECT_FOLDER}/storage/datasets/default/ 에 JSON 파일로 저장됩니다.

- custom Dataset을 만들고 싶다면 Dataset.open() 함수를 사용할 수 있습니다.

```typescript
import { Dataset, CheerioCrawler } from "crawlee";

const crawler = new CheerioCrawler({
  // Function called for each URL
  async requestHandler({ request, body }) {
    // Save data to default dataset
    await Dataset.pushData({
      url: request.url,
      html: body,
    });
  },
});

await crawler.addRequests([
  "http://www.example.com/page-1",
  "http://www.example.com/page-2",
  "http://www.example.com/page-3",
]);

// Run the crawler
await crawler.run();
```

## Examples

[참조](https://crawlee.dev/docs/examples/accept-user-input)

- Playwright crawler sample

```typescript
import { Dataset, PlaywrightCrawler } from "crawlee";

// Create an instance of the PlaywrightCrawler class - a crawler
// that automatically loads the URLs in headless Chrome / Playwright.
const crawler = new PlaywrightCrawler({
  launchContext: {
    // Here you can set options that are passed to the playwright .launch() function.
    launchOptions: {
      headless: true,
    },
  },

  // Stop crawling after several pages
  maxRequestsPerCrawl: 50,

  // This function will be called for each URL to crawl.
  // Here you can write the Playwright scripts you are familiar with,
  // with the exception that browsers and pages are automatically managed by Crawlee.
  // The function accepts a single parameter, which is an object with a lot of properties,
  // the most important being:
  // - request: an instance of the Request class with information such as URL and HTTP method
  // - page: Playwright's Page object (see https://playwright.dev/docs/api/class-page)
  async requestHandler({ request, page, enqueueLinks, log }) {
    log.info(`Processing ${request.url}...`);

    // A function to be evaluated by Playwright within the browser context.
    const data = await page.$$eval(".athing", ($posts) => {
      const scrapedData: { title: string; rank: string; href: string }[] = [];

      // We're getting the title, rank and URL of each post on Hacker News.
      $posts.forEach(($post) => {
        scrapedData.push({
          title: $post.querySelector(".title a").innerText,
          rank: $post.querySelector(".rank").innerText,
          href: $post.querySelector(".title a").href,
        });
      });

      return scrapedData;
    });

    // Store the results to the default dataset.
    await Dataset.pushData(data);

    // Find a link to the next page and enqueue it if it exists.
    const infos = await enqueueLinks({
      selector: ".morelink",
    });

    if (infos.processedRequests.length === 0)
      log.info(`${request.url} is the last page!`);
  },

  // This function is called if the page processing failed more than maxRequestRetries+1 times.
  failedRequestHandler({ request, log }) {
    log.info(`Request ${request.url} failed too many times.`);
  },
});

await crawler.addRequests(["https://news.ycombinator.com/"]);

// Run the crawler and wait for it to finish.
await crawler.run();

console.log("Crawler finished.");
```
